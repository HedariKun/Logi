package logi

import "runtime"

func GetCallerInfo(skip int) (string, int, bool) {
	_, file, line, ok := runtime.Caller(skip)
	return file, line, ok
}
