package logi

import (
	"fmt"
	"time"
)

type LogLevel string

const (
	InfoLevel  LogLevel = "INFO"
	DebugLevel LogLevel = "DEBUG"
	WarnLevel  LogLevel = "WARN"
	ErrorLevel LogLevel = "ERROR"
	FatalLevel LogLevel = "FATAL"
)

type Logger interface {
	Info(entry Entry)
	Debug(meesage Entry)
	Warn(meesage Entry)
	Error(message Entry)
	Fatal(message Entry)
}

type Manager struct {
	Loggers []Logger

	options ManagerOptions

	context map[string]string
}

type ManagerOptions struct {
	ShowTimeStamp bool `json:"show_time_stamp"`

	ShowFileName   bool `json:"show_file_name"`
	ShowLineNumber bool `json:"show_line_number"`

	ShowFields bool `json:"show_fields"`

	ShowInfoLevel  bool `json:"show_info"`
	ShowDebugLevel bool `json:"show_debug"`
	ShowWarnLevel  bool `json:"show_warn"`
	ShowErrorLevel bool `json:"show_error"`
	ShowFatalLevel bool `json:"show_fatal"`
}

func New(options ...ManagerOptions) *Manager {
	manager := &Manager{}

	if len(options) <= 0 {
		manager.options = ManagerOptions{
			ShowTimeStamp:  true,
			ShowFileName:   true,
			ShowLineNumber: true,

			ShowFields: true,

			ShowInfoLevel:  true,
			ShowDebugLevel: true,
			ShowWarnLevel:  true,
			ShowErrorLevel: true,
			ShowFatalLevel: true,
		}
	} else {
		manager.options = options[0]
	}

	manager.context = make(map[string]string)
	return manager
}

func (m *Manager) AddLogger(logger Logger) *Manager {
	m.Loggers = append(m.Loggers, logger)
	return m
}

func (m *Manager) With(key string, value any) *Manager {
	newManager := *m

	newManager.context = make(map[string]string)
	for k, v := range m.context {
		newManager.context[k] = v
	}
	newManager.context[key] = fmt.Sprintf("%+v", value)

	return &newManager
}

func (m *Manager) PrepareEntry(level LogLevel, message string) *Entry {
	entry := NewEntry()
	entry.manager = m
	entry.Level = level
	entry.Message = message
	entry.fmt = "$Time$ [$Level$] $File$ $Field$: $Message$\n"

	if m.options.ShowFields {
		entry.Fields = m.context
	}

	if m.options.ShowTimeStamp {
		entry.TimeStamp = time.Now()
	}

	file, line, ok := GetCallerInfo(3)

	if ok {
		if m.options.ShowFileName {
			entry.FileName = file
		}
		if m.options.ShowLineNumber {
			entry.LineNumber = line
		}
	}

	if m.options.ShowFields {
		entry.Fields = m.context
	}

	return entry
}

func (m *Manager) Info(message string) {
	if !m.options.ShowInfoLevel {
		return
	}

	entry := m.PrepareEntry(InfoLevel, message)

	for _, logger := range m.Loggers {
		logger.Info(*entry)
	}
}

func (m *Manager) Debug(message string) {
	if !m.options.ShowDebugLevel {
		return
	}

	entry := m.PrepareEntry(DebugLevel, message)

	for _, logger := range m.Loggers {
		logger.Debug(*entry)
	}
}

func (m *Manager) Warn(message string) {
	if !m.options.ShowWarnLevel {
		return
	}

	entry := m.PrepareEntry(WarnLevel, message)

	for _, logger := range m.Loggers {
		logger.Warn(*entry)
	}
}

func (m *Manager) Error(message string) {
	if !m.options.ShowErrorLevel {
		return
	}

	entry := m.PrepareEntry(ErrorLevel, message)

	for _, logger := range m.Loggers {
		logger.Error(*entry)
	}
}

func (m *Manager) Fatal(message string) {
	if !m.options.ShowFatalLevel {
		return
	}

	entry := m.PrepareEntry(FatalLevel, message)

	for _, logger := range m.Loggers {
		logger.Fatal(*entry)
	}
}
