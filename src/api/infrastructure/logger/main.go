package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

func fileEncoder(encoderConfig *zapcore.EncoderConfig) (zapcore.Encoder, zapcore.WriteSyncer) {

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	wd, _ := os.Getwd()
	logFile, _ := os.OpenFile(filepath.Join(wd, "/test.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	encoder := zapcore.NewJSONEncoder(*encoderConfig)
	writeSyncer := zapcore.AddSync(logFile)

	return encoder, writeSyncer
}

func consoleEncoder(encoderConfig *zapcore.EncoderConfig) (zapcore.Encoder, zapcore.WriteSyncer) {
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(*encoderConfig)
	writeSyncer := zapcore.AddSync(os.Stdout)

	return encoder, writeSyncer
}

func UseZap() *zap.Logger {
	encoderCfg := &zapcore.EncoderConfig{
		TimeKey:          "timestamp",
		LevelKey:         "level",
		FunctionKey:      "func",
		CallerKey:        "caller",
		MessageKey:       "msg",
		StacktraceKey:    zapcore.OmitKey,
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeTime:       syslogTimeEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: "  |  ",
	}

	fEncoder, fileWriter := fileEncoder(encoderCfg)
	cEncoder, consoleWriter := consoleEncoder(encoderCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(fEncoder, fileWriter, zapcore.DebugLevel),
		zapcore.NewCore(cEncoder, consoleWriter, zapcore.DebugLevel),
	)

	return zap.New(core)
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
