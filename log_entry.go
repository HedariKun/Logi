package logi

import (
	"fmt"
	"strings"
	"time"
)

type EntryTag string

const (
	EntryTagMessage EntryTag = "Message"
	EntryTagLevel   EntryTag = "Level"
	EntryTagTime    EntryTag = "Time"
	EntryTagFile    EntryTag = "File"
	EntryTagField   EntryTag = "Field"
)

type Entry struct {
	manager *Manager

	Level     LogLevel  `json:"level"`
	Message   string    `json:"message"`
	TimeStamp time.Time `json:"time_stamp"`

	fmt string

	FileName   string `json:"file_name"`
	LineNumber int    `json:"line_number"`

	Fields map[string]string `json:"meta_data"`

	hooks map[string]func(*Entry) string
}

func NewEntry() *Entry {
	entry := Entry{}
	entry.hooks = make(map[string]func(*Entry) string)

	entry.hooks[string(EntryTagMessage)] = func(e *Entry) string { return e.Message }
	entry.hooks[string(EntryTagLevel)] = func(e *Entry) string { return string(e.Level) }
	entry.hooks[string(EntryTagTime)] = func(e *Entry) string {
		if e.manager.options.ShowTimeStamp {
			return e.TimeStamp.Format(time.RFC3339)
		} else {
			return ""
		}
	}
	entry.hooks[string(EntryTagFile)] = func(e *Entry) string {
		if e.FileName != "" && e.LineNumber != 0 {
			return e.FileName + ":" + fmt.Sprintf("%d", e.LineNumber)
		} else if e.FileName != "" && e.LineNumber == 0 {
			return e.FileName
		} else if e.FileName == "" && e.LineNumber != 0 {
			return fmt.Sprintf("%d", e.LineNumber)
		} else {
			return ""
		}
	}

	entry.hooks[string(EntryTagField)] = func(e *Entry) string {
		if e.manager.options.ShowFields && len(e.Fields) > 0 {
			value := ""
			for k, v := range e.Fields {
				value += fmt.Sprintf("%s=%s ", k, v)
			}
			return strings.TrimSpace(value)
		} else {
			return ""
		}
	}

	return &entry
}

func (e *Entry) BeforeTag(tag EntryTag, hook func(*Entry) string) {
	e.hooks["before:"+string(tag)] = hook
}

func (e *Entry) AfterTag(tag EntryTag, hook func(*Entry) string) {
	e.hooks["after:"+string(tag)] = hook
}

func (e *Entry) BeforeAll(hook func(*Entry) string) {
	e.hooks["before:all"] = hook
}

func (e *Entry) AfterAll(hook func(*Entry) string) {
	e.hooks["after:all"] = hook
}

func (e *Entry) HandleTag(tag EntryTag, hook func(*Entry) string) {
	e.hooks[string(tag)] = hook
}

func (e *Entry) String() string {
	res := e.fmt

	tags := []EntryTag{EntryTagMessage, EntryTagLevel, EntryTagTime, EntryTagFile, EntryTagField}

	beforeAll, beforeAllOk := e.hooks["before:all"]
	afterAll, afterAllOk := e.hooks["after:all"]

	for _, tag := range tags {
		tagValue := ""

		tagHook, ok := e.hooks[string(tag)]
		if !ok {
			continue
		}

		tagValue += tagHook(e)

		beforeHook, ok := e.hooks["before:"+string(tag)]
		if ok {
			tagValue = beforeHook(e) + tagValue
		}

		if beforeAllOk {
			tagValue = beforeAll(e) + tagValue
		}

		afterHook, ok := e.hooks["after:"+string(tag)]
		if ok {
			tagValue += afterHook(e)
		}

		if afterAllOk {
			tagValue += afterAll(e)
		}

		res = strings.Replace(res, "$"+string(tag)+"$", tagValue, 1)
	}

	return res

}
