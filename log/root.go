package log

import (
	"os"
)

var rootLogger *Logger

func init() {
	rootLogger = &Logger{}
	rootLogger.AddAppender(NewWriterAppender(os.Stdout, NewSimpleLayout()))
	rootLogger.threshold = PriorityInfo

	logEnv := os.Getenv("LOG")
	if logEnv != "" {
		threshold, ok := ParsePriority(logEnv)
		if !ok {
			Warn("Unable to parse priority level", logEnv)
		} else {
			rootLogger.threshold = threshold
		}
	}
}

func Fatal(msg string, params ...interface{}) {
	rootLogger.Fatal(msg, params...)
}

func Error(msg string, params ...interface{}) {
	rootLogger.Error(msg, params...)
}

func Warn(msg string, params ...interface{}) {
	rootLogger.Warn(msg, params...)
}

func Info(msg string, params ...interface{}) {
	rootLogger.Info(msg, params...)
}

func Debug(msg string, params ...interface{}) {
	rootLogger.Debug(msg, params...)
}
