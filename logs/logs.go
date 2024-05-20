package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	// Log, _ = zap.NewProduction()
	// zap.NewProduction (json) / zap.NewDevelopment (console) is zap configuration preset

	// custom config
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // YYYY-MM-DDThh:mm:ss[.mmm]TZD
	config.EncoderConfig.StacktraceKey = ""                      // disable stacktrace

	var err error
	log, err = config.Build(zap.AddCallerSkip(1)) // skip 1 step from caller
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

func Error(message interface{}, fields ...zap.Field) {
	/*
		msg, ok := message.(error)
		if ok {
			log.Error(msg.Error(), fields...)
		}
	*/

	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
