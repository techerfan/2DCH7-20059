package zap

import (
	"github.com/techerfan/2DCH7-20059/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type log struct {
	zap *zap.SugaredLogger
}

// Creates a new logger. max size is in megabytes. if you set max backup and max age to 0,
// no old log files will be deleted.
func New(
	path string,
	maxSize int,
	maxBackups int,
	maxAge int,
	level zapcore.Level,
) (logger.Logger, error) {
	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	})

	encodingConfig := zap.NewProductionEncoderConfig()
	encodingConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodingConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	enc := zapcore.NewJSONEncoder(encodingConfig)
	core := zapcore.NewCore(enc, ws, level)

	z := zap.New(core)
	sugared := z.Sugar()

	return &log{sugared}, nil
}

func (l *log) Sync() {
	l.zap.Sync()
}

func (l *log) Error(msg string, kv ...interface{}) {
	l.zap.Errorw(msg, kv...)
}

func (l *log) Errorf(msg string, args ...interface{}) {
	l.zap.Errorf(msg, args...)
}

func (l *log) Warn(msg string, kv ...interface{}) {
	l.zap.Warnw(msg, kv...)
}

func (l *log) Warnf(msg string, args ...interface{}) {
	l.zap.Warnf(msg, args...)
}

func (l *log) Info(msg string, kv ...interface{}) {
	l.zap.Infow(msg, kv...)
}

func (l *log) Infof(msg string, args ...interface{}) {
	l.zap.Infof(msg, args...)
}
