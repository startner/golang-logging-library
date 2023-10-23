package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type ZapField struct {
	Key   string
	Value any
}

const ERROR = "ERROR"
const WARN = "WARN"
const INFO = "INFO"
const DEBUG = "DEBUG"

func LogCustom(logLevel string, msg string, functionName string, data string, fields ...zap.Field) {
	allFields := make([]zap.Field, 0)

	// Add common fields
	allFields = append(allFields, zap.String("uuid", UuidLogger))
	allFields = append(allFields, zap.String("functionName", functionName))
	allFields = append(allFields, zap.String("data", data))

	// Append any additional fields
	allFields = append(allFields, fields...)

	switch logLevel {
	case ERROR:
		Logger.Error(fmt.Sprintf("%v", msg), allFields...)
	case WARN:
		Logger.Warn(fmt.Sprintf("%v", msg), allFields...)
	case INFO:
		Logger.Info(fmt.Sprintf("%v", msg), allFields...)
	case DEBUG:
		Logger.Debug(fmt.Sprintf("%v", msg), allFields...)
	default:
		Logger.Info(fmt.Sprintf("%v", msg), allFields...)
	}
}
