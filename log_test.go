package bench_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/exp/slog"

	gokitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func BenchmarkLog_LogrusDebugLevel(b *testing.B) {
	logger, cleanUp := setupLogrusLogger()
	defer cleanUp()

	b.ResetTimer()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		logger.WithFields(logrus.Fields{
			"name": "Nuruddin Ashr",
		}).Info("Hello, World!")

		logger.WithFields(logrus.Fields{
			"name": "Nuruddin Ashr",
		}).Debug("Hello, World!")
	}
}
func BenchmarkLog_LogrusInfoLevel(b *testing.B) {
	logger, cleanUp := setupLogrusLogger()
	defer cleanUp()

	logger.SetLevel(logrus.InfoLevel)

	b.ResetTimer()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		logger.WithFields(logrus.Fields{
			"name": "Nuruddin Ashr",
		}).Info("Hello, World!")

		logger.WithFields(logrus.Fields{
			"name": "Nuruddin Ashr",
		}).Debug("Hello, World!")
	}
}

func BenchmarkLog_GokitDebugLevel(b *testing.B) {
	logger, cleanUp := setupGokitLogger()
	defer cleanUp()

	b.ResetTimer()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		level.Info(logger).Log(
			"msg", "Hello, World!",
			"name", "Nuruddin Ashr",
		)

		level.Debug(logger).Log(
			"msg", "Hello, World!",
			"name", "Nuruddin Ashr",
		)
	}
}

func BenchmarkLog_GokitInfoLevel(b *testing.B) {
	logger, cleanUp := setupGokitLogger()
	defer cleanUp()

	logger = level.NewFilter(logger, level.AllowInfo())

	b.ResetTimer()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		level.Info(logger).Log(
			"msg", "Hello, World!",
			"name", "Nuruddin Ashr",
		)

		level.Debug(logger).Log(
			"msg", "Hello, World!",
			"name", "Nuruddin Ashr",
		)
	}
}

func BenchmarkLog_SlogDebugLevel(b *testing.B) {
	logger, cleanUp := setupSlogLogger(slog.HandlerOptions{Level: slog.LevelDebug})
	defer cleanUp()

	b.ResetTimer()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("Hello, World!", slog.String("name", "Nuruddin Ashr"))
		logger.Debug("Hello, World!", slog.String("name", "Nuruddin Ashr"))
	}
}

func BenchmarkLog_SlogInfoLevel(b *testing.B) {
	logger, cleanUp := setupSlogLogger(slog.HandlerOptions{Level: slog.LevelInfo})
	defer cleanUp()

	b.ResetTimer()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("Hello, World!", slog.String("name", "Nuruddin Ashr"))
		logger.Debug("Hello, World!", slog.String("name", "Nuruddin Ashr"))
	}
}

func BenchmarkLog_ZapDebugLevel(b *testing.B) {
	logger, cleanUp := setupZapLogger(zap.DebugLevel)
	defer cleanUp()

	b.ResetTimer()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("Hello, World!", zap.String("name", "Nuruddin Ashr"))
		logger.Debug("Hello, World!", zap.String("name", "Nuruddin Ashr"))
		// logger.Sync()
	}
}

func BenchmarkLog_ZapInfoLevel(b *testing.B) {
	logger, cleanUp := setupZapLogger(zap.InfoLevel)
	defer cleanUp()

	b.ResetTimer()
	defer b.StopTimer()

	for i := 0; i < b.N; i++ {
		logger.Info("Hello, World!", zap.String("name", "Nuruddin Ashr"))
		logger.Debug("Hello, World!", zap.String("name", "Nuruddin Ashr"))
		// logger.Sync()
	}
}

func setupLogrusLogger() (*logrus.Logger, func()) {
	f, err := os.CreateTemp("", "logbench.*.log")
	if err != nil {
		panic(fmt.Sprint("Fail to create temp file:", err))
	}

	logger := logrus.New()
	logger.SetOutput(f)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	cleanUp := func() {
		f.Close()
		os.Remove(f.Name())
	}
	return logger, cleanUp
}

func setupGokitLogger() (gokitlog.Logger, func()) {
	f, err := os.CreateTemp("", "logbench.*.log")
	if err != nil {
		panic(fmt.Sprint("Fail to create temp file:", err))
	}

	logger := gokitlog.NewJSONLogger(f)
	gokitlog.With(logger, "time", gokitlog.TimestampFormat(time.Now, time.RFC3339))

	cleanUp := func() {
		f.Close()
		os.Remove(f.Name())
	}
	return logger, cleanUp
}

func setupSlogLogger(opts slog.HandlerOptions) (*slog.Logger, func()) {
	f, err := os.CreateTemp("", "logbench.*.log")
	if err != nil {
		panic(fmt.Sprint("Fail to create temp file:", err))
	}

	handler := opts.NewJSONHandler(f)

	cleanUp := func() {
		f.Close()
		os.Remove(f.Name())
	}

	logger := slog.New(handler)
	return logger, cleanUp
}

func setupZapLogger(level zapcore.Level) (*zap.Logger, func()) {
	f, err := os.CreateTemp("", "logbench.*.log")
	if err != nil {
		panic(fmt.Sprint("Fail to create temp file:", err))
	}

	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.ISO8601TimeEncoder

	jsonEncoder := zapcore.NewJSONEncoder(ec)
	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(f), level),
	)

	logger := zap.New(core)

	cleanUp := func() {
		f.Close()
		os.Remove(f.Name())
	}

	return logger, cleanUp
}
