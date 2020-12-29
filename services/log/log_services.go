package log

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

const DefaultLoggerBufferSize = 100

var GLogger *Logger

type Logger struct {
	logs.BeeLogger
}

func (logger *Logger) setFileLogger(fileName string, separateFile bool, level int) error {
	if separateFile {
		config := fmt.Sprintf(`{"filename":"%s","level":%d,"separate":["error"]}`, fileName, level)
		return logger.SetLogger(logs.AdapterMultiFile, config)
	}
	config := fmt.Sprintf(`{"filename":"%s","level":%d}`, fileName, level)
	return logger.SetLogger(logs.AdapterFile, config)
}

func (logger *Logger) setConsoleLogger(level int) error {
	config := fmt.Sprintf(`{"level":%d}`, level)
	return logger.SetLogger(logs.AdapterConsole, config)
}

func NewLogger(fileName string) (*Logger, error) {
	logger := &Logger{*logs.NewLogger(DefaultLoggerBufferSize)}
	err := logger.setFileLogger(fileName, true, logs.LevelInfo)
	if err != nil {
		return nil, err
	}
	err = logger.setConsoleLogger(logs.LevelCritical)
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func SetLogger(logger *Logger) {
	GLogger = logger
}
