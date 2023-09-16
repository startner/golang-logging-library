package logger

import (
	"fmt"
	"go.uber.org/zap"
)

type ZapField struct {
	Key   string
	Value any
}

func LogError(str any) {
	Logger.Error(fmt.Sprintf("%v", str),
		zap.String("uuid", UuidLogger),
	)
}

func LogErrorWithFunctionName(str any, functionName string) {
	Logger.Error(fmt.Sprintf("%v", str),
		zap.String("uuid", UuidLogger),
		zap.String("functionName", functionName),
	)
}

func LogInfo(str any) {
	Logger.Info(fmt.Sprintf("%v", str),
		zap.String("uuid", UuidLogger),
	)
}

func LogInfoWithFunctionName(str any, functionName string) {
	Logger.Info(fmt.Sprintf("%v", str),
		zap.String("uuid", UuidLogger),
		zap.String("functionName", functionName),
	)
}

func LogWarn(str any) {
	Logger.Warn(fmt.Sprintf("%v", str),
		zap.String("uuid", UuidLogger),
	)
}

func LogWarnWithFunctionName(str any, functionName string) {
	Logger.Warn(fmt.Sprintf("%v", str),
		zap.String("uuid", UuidLogger),
		zap.String("functionName", functionName),
	)
}
