package logi

type Logger interface {
	Log(string)
	Error(string)
	Warning(string)
}

type LogType string

const (
	TYPE_LOG     LogType = "Log"
	TYPE_ERROR   LogType = "Error"
	TYPE_WARNING LogType = "Warning"
)

type LoggerManager struct {
	loggers   []Logger
	formatter Formater
}

func New() *LoggerManager {
	manager := LoggerManager{}
	manager.formatter = BasicFormater{
		ShowTime: true,
	}
	return &manager
}

func (manager *LoggerManager) Add(logger Logger) *LoggerManager {
	manager.loggers = append(manager.loggers, logger)
	return manager
}

func (manager LoggerManager) Log(message string) {
	msg := manager.formatter.Format(message, TYPE_LOG)
	for _, logger := range manager.loggers {
		logger.Log(msg)
	}
}

func (manager LoggerManager) Error(message string) {
	msg := manager.formatter.Format(message, TYPE_ERROR)
	for _, logger := range manager.loggers {
		logger.Error(msg)
	}
}

func (manager LoggerManager) Warning(message string) {
	msg := manager.formatter.Format(message, TYPE_WARNING)
	for _, logger := range manager.loggers {
		logger.Warning(msg)
	}
}
