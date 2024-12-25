package logger

import (
    "fmt"
    "os"
    "path/filepath"
    "time"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"

)

var log *zap.Logger

func Init(logFile string, level string) error {
    // 创建日志目录
    if err := os.MkdirAll(filepath.Dir(logFile), 0744); err != nil {
        return fmt.Errorf("创建日志目录失败: %v", err)
    }

    // 设置日志级别
    var logLevel zapcore.Level
    switch level {
    case "debug":
        logLevel = zapcore.DebugLevel
    case "info":
        logLevel = zapcore.InfoLevel
    case "warn":
        logLevel = zapcore.WarnLevel
    case "error":
        logLevel = zapcore.ErrorLevel
    default:
        logLevel = zapcore.InfoLevel
    }

    // 配置 zapcore
    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     timeEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }

    // 配置日志输出和切割
    writer := zapcore.AddSync(&lumberjack.Logger{
        Filename:   logFile,
        MaxSize:    100, // MB
        MaxBackups: 3,
        MaxAge:     28,   // days
        Compress:   true, // 是否压缩
    })

    // 同时输出到控制台和文件
    core := zapcore.NewTee(
        zapcore.NewCore(
            zapcore.NewJSONEncoder(encoderConfig),
            writer,
            logLevel,
        ),
        zapcore.NewCore(
            zapcore.NewConsoleEncoder(encoderConfig),
            zapcore.AddSync(os.Stdout),
            logLevel,
        ),
    )

    // 创建 logger
    log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
    return nil
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 导出日志方法
func Debug(msg string, fields ...zap.Field) {
    log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
    log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
    log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
    log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
    log.Fatal(msg, fields...)
} 