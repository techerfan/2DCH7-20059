package logger

import (
	"log"
)

//go:generate mockgen -destination=./mock.go -package=logger . Logger
type Logger interface {
	Info(msg string, kv ...interface{})
	Infof(msg string, args ...interface{})
	Warn(msg string, kv ...interface{})
	Warnf(msg string, args ...interface{})
	Error(msg string, kv ...interface{})
	Errorf(msg string, args ...interface{})
	Sync()
}

type DummyLogger struct{}

func (d DummyLogger) Info(msg string, kv ...interface{}) {
	log.Println(msg, kv)
}

func (d DummyLogger) Infof(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func (d DummyLogger) Warn(msg string, kv ...interface{}) {
	log.Println("[WARN] ", msg, kv)
}

func (d DummyLogger) Warnf(msg string, args ...interface{}) {
	log.Printf("[WARN] "+msg, args)

}

func (d DummyLogger) Error(msg string, kv ...interface{}) {
	log.Println("[ERROR] ", msg, kv)

}

func (d DummyLogger) Errorf(msg string, args ...interface{}) {
	log.Printf("[ERROR] "+msg, args)
}

func (d DummyLogger) Sync() {

}
