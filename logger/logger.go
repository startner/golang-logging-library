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

func LogErrorCustom(str any, fields ...zap.Field) {
	Logger.Error(fmt.Sprintf("%v", str), fields...)
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
func LogInfoCustom(str any, fields ...zap.Field) {
	Logger.Error(fmt.Sprintf("%v", str), fields...)
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
func LogWarnCustom(str any, fields ...zap.Field) {
	Logger.Error(fmt.Sprintf("%v", str), fields...)
}

func LogDebug(str any) {
	Logger.Warn(fmt.Sprintf("%v", str),
		zap.String("uuid", UuidLogger),
	)
}

func LogDebugWithFunctionName(str any, functionName string) {
	Logger.Warn(fmt.Sprintf("%v", str),
		zap.String("uuid", UuidLogger),
		zap.String("functionName", functionName),
	)
}
func LogDebugCustom(str any, fields ...zap.Field) {
	Logger.Error(fmt.Sprintf("%v", str), fields...)
}
