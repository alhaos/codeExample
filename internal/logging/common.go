package logging

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	zapSugarLogger *zap.SugaredLogger
}

type Config struct {
	Directory  string `yaml:"directory" validate:"require"`
	Filename   string `yaml:"filename" validate:"require"`
	MaxSize    int    `yaml:"maxSize" validate:"require"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
	Level      string `yaml:"level"`
}

func New(config *Config) *Logger {

	stdoutSyncer := zapcore.AddSync(os.Stdout)

	lumberjackWriter := lumberjack.Logger{
		Filename:   filepath.Join(config.Directory, config.Filename),
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
	}

	level := zap.NewAtomicLevel()

	err := level.UnmarshalText([]byte(config.Level))
	if err != nil {
		CrashLog(fmt.Errorf("unable to unmarshal logging level: %w", err))
	}

	fileSyncer := zapcore.AddSync(&lumberjackWriter)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(time.Stamp))
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, stdoutSyncer, level),
		zapcore.NewCore(encoder, fileSyncer, level),
	)

	logger := &Logger{zapSugarLogger: zap.New(core).Sugar()}

	return logger
}

func (l Logger) Debug(args ...interface{}) {
	l.zapSugarLogger.Debug(args...)
}

func (l Logger) Debugf(template string, args ...interface{}) {
	l.zapSugarLogger.Debugf(template, args...)
}

func (l Logger) Info(args ...interface{}) {
	l.zapSugarLogger.Info(args...)
}

func (l Logger) Infof(template string, args ...interface{}) {
	l.zapSugarLogger.Infof(template, args...)
}

func (l Logger) Error(args ...interface{}) {
	l.zapSugarLogger.Error(args...)
}

func (l Logger) Errorf(template string, args ...interface{}) {
	l.zapSugarLogger.Errorf(template, args...)
}

func (l Logger) Fatal(args ...interface{}) {
	l.zapSugarLogger.Fatal(args...)
}

func (l Logger) Fatalf(template string, args ...interface{}) {
	l.zapSugarLogger.Fatalf(template, args...)
}

func CrashLog(e error) {

	crashLogFileName := fmt.Sprintf("crash-%s.log", time.Now().Format("2006-01-02T15-04-05"))
	f, err := os.OpenFile(crashLogFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	line := fmt.Sprintf("%s: %s", time.Now().Format(time.DateTime), e.Error())

	_, err = f.WriteString(line)
	if err != nil {
		panic(err)
	}
}
