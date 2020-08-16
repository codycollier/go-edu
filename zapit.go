package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	// default of std lib log: 2020/08/15 21:44:07 foo
	// defaultish console zap: 2020-08-15T21:44:07.012-0500	INFO	foo

	// Explore console style logging with zap

	zEncoderConfig := zapcore.EncoderConfig{
		TimeKey:          "yes",
		LevelKey:         "yes",
		NameKey:          "yes",
		CallerKey:        "yes",
		FunctionKey:      zapcore.OmitKey,
		MessageKey:       "yes",
		StacktraceKey:    "yes",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " ",
	}

	zConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     zEncoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}

	zlog, err := zConfig.Build()
	if err != nil {
		println("error creating zap logger")
	}
	defer zlog.Sync()

	for i := 0; i < 5; i++ {
		zlog.Info("foo", zap.Int("count", i))
		zlog.Debug("baz", zap.Int("count", i))
	}

	zlogs := zlog.Sugar()
	for i := 0; i < 5; i++ {
		zlogs.Infof("foo count: %v", i)
		zlogs.Warnf("baz count: %v", i)
	}
	zlog.Error("oh no!")
	zlog.Info("back on track!")
	zlog.Info("done")

}
