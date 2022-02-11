package standardLogger

import (
	"fmt"
	"log"
	"os"

	"github.com/kazmerdome/godome/pkg/observer/logger"
)

type standardLogger struct {
	logger *log.Logger
}

func NewStandardLogger() logger.Logger {
	logger := log.New(os.Stdout, "", 0)
	return &standardLogger{logger}
}

func (r *standardLogger) Info(msg string, a ...interface{}) {
	r.logger.Printf(msg, a...)
}

func (r *standardLogger) Fatal(msg string, a ...interface{}) {
	r.logger.Fatal(fmt.Sprintf(msg, a...))

}

func (r *standardLogger) Warn(msg string, a ...interface{}) {
	r.logger.Printf(msg, a...)
}

func (r *standardLogger) Error(msg string, a ...interface{}) {
	r.logger.Fatal(fmt.Sprintf(msg, a...))
}

func (r *standardLogger) Panic(msg string, a ...interface{}) {
	r.logger.Panic(fmt.Sprintf(msg, a...))
}
