package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//ZapLogger contains the reference to zap's log level
type ZapLogger struct {
	Level *zap.AtomicLevel
	Log   *zap.Logger
}

//NewLogger creates a new Uber zap logger
func NewLogger() *ZapLogger {
	atom := zap.NewAtomicLevelAt(zapcore.DebugLevel)

	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       atom,
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()

	return &ZapLogger{
		Level: &atom,
		Log:   logger,
	}
}

//UpdateLevel changes the zap log level
func UpdateLevel(lvl string, atom *zap.AtomicLevel) {
	zaplvl := getZapLevel(lvl)
	atom.SetLevel(zaplvl)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "Info":
		return zapcore.InfoLevel
	case "Warn":
		return zapcore.WarnLevel
	case "Debug":
		return zapcore.DebugLevel
	case "Error":
		return zapcore.ErrorLevel
	case "Fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
