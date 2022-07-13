package logger

import (
	"fmt"
	"os"
	"simpson/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

type ILogger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

type logger struct {
	sugarLogger *zap.SugaredLogger
	*zap.Logger
}

var Logger *logger = &logger{}

func configure() zapcore.WriteSyncer {
	writers := []zapcore.WriteSyncer{os.Stderr}
	return zapcore.NewMultiWriteSyncer(writers...)
}

// App Logger constructor
func Newlogger(lifecycle fx.Lifecycle, config *config.Config, mode, level, format string) ILogger {
	// lifecycle.Append(fx.Hook{
	// 	OnStart: func(context.Context) error {
	// 		return nil
	// 	},
	// 	OnStop: func(ctx context.Context) error {
	// 		return nil
	// 	},
	// })

	fmt.Println("di-new logger")
	// lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
	// 	fmt.Println("config", config.Logger.Level)
	// 	return nil
	// }})
	logLevel, exist := loggerLevelMap[level]
	if !exist {
		logLevel = zapcore.DebugLevel
	}

	var encoderCfg zapcore.EncoderConfig
	if mode == "pro" {
		encoderCfg = zap.NewProductionEncoderConfig()
	} else {
		encoderCfg = zap.NewDevelopmentEncoderConfig()

	}
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"
	encoderCfg.EncodeDuration = zapcore.NanosDurationEncoder
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if format == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}
	core := zapcore.NewCore(encoder, configure(), zap.NewAtomicLevelAt(logLevel))
	loggerzap := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	sugarLogger := loggerzap.Sugar()
	// Logger := &logger{
	// 	sugarLogger: sugarLogger,
	// 	Logger:      loggerzap,
	// }
	// return *Logger
	tempLog := &logger{
		sugarLogger: sugarLogger,
		Logger:      loggerzap,
	}
	Logger = tempLog
	return tempLog
}

func (l *logger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *logger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
