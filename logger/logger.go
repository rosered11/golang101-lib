package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoding := zap.NewProductionEncoderConfig()
	encoding.TimeKey = "timestamp"
	encoding.EncodeTime = zapcore.ISO8601TimeEncoder

	config.EncoderConfig = encoding

	//log, err = zap.NewProduction(zap.AddCallerSkip(1))
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
