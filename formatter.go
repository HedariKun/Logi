package logi

import (
	"fmt"
	"time"
)

type Formater interface {
	Format(message string, typ LogType) string
}

type BasicFormater struct {
	ShowTime bool
}

func (formater BasicFormater) Format(message string, typ LogType) string {
	val := ""

	val += "[" + string(typ) + "]"

	if formater.ShowTime {
		date := time.Now()
		formatedDate := date.Format("2006-01-02-15-04-05:000")
		val += fmt.Sprintf("[%s]", formatedDate)
	}
	val += ": " + message

	return val
}
