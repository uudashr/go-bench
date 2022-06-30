package bench_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	kitlog "github.com/go-kit/log"
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
	logger, cleanUp := setupKitLogger()
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
	logger, cleanUp := setupKitLogger()
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

func setupLogrusLogger() (*logrus.Logger, func()) {
	f, err := ioutil.TempFile("", "logbench.*.log")
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

func setupKitLogger() (kitlog.Logger, func()) {
	f, err := ioutil.TempFile("", "logbench.*.log")
	if err != nil {
		panic(fmt.Sprint("Fail to create temp file:", err))
	}

	logger := kitlog.NewJSONLogger(f)
	kitlog.With(logger, "time", kitlog.TimestampFormat(time.Now, time.RFC3339))

	cleanUp := func() {
		f.Close()
		os.Remove(f.Name())
	}
	return logger, cleanUp
}
