package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.Logger
var UuidLogger string

func InitLogger(timezone string) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		customTimeEncoder(t, enc, timezone)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		zap.InfoLevel,
	)

	Logger = zap.New(core)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder, timezone string) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.UTC
	}

	t = t.In(loc)
	enc.AppendString(t.Format("20060102 15:04:05"))
}
