package log

import (
	"fmt"
	"github.com/smallnest/rpcx/log"
	"os"
	"runtime"
	"strconv"
)


type Logger interface {
	log.Logger
}

var defaultLogger Logger = &stdLogger{}


func Debug(v ...interface{}) {
	defaultLogger.Debug(v...)
}

func Debugf(fmt string, v ...interface{}) {
	defaultLogger.Debugf(fmt, v...)
}

func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

func Infof(fmt string, v ...interface{}) {
	defaultLogger.Infof(fmt, v...)
}

func  Warn(v ...interface{}) {
	defaultLogger.Warn(v...)
}

func  Warnf(fmt string, v ...interface{}) {
	defaultLogger.Warnf(fmt, v...)
}

func  Error(v ...interface{}) {
	defaultLogger.Error(v...)
}

func  Errorf(fmt string, v ...interface{}) {
	defaultLogger.Errorf(fmt, v...)
}

func  Fatal(v ...interface{}) {
	defaultLogger.Fatal(v...)
}

func  Fatalf(fmt string, v ...interface{}) {
	defaultLogger.Fatalf(fmt, v...)
}

func Panic(v ...interface{}) {
	defaultLogger.Panic(v...)
}

func Panicf(fmt string, v ...interface{}) {
	defaultLogger.Panicf(fmt, v...)
}


func SetLogger(logger Logger) {
	defaultLogger = logger
	log.SetLogger(logger)
}

type stdLogger struct {

}

func (dl *stdLogger) Debug(v ...interface{}) {
	dl.output(v...)
}

func (dl *stdLogger) Debugf(fmt string, v ...interface{}) {
	dl.outputf(fmt, v...)
}

func (dl *stdLogger) Info(v ...interface{}) {
	dl.output(v...)
}

func (dl *stdLogger) Infof(fmt string, v ...interface{}) {
	dl.outputf(fmt, v...)
}

func (dl *stdLogger) Warn(v ...interface{}) {
	dl.output(v...)
}

func (dl *stdLogger) Warnf(fmt string, v ...interface{}) {
	dl.outputf(fmt, v...)
}

func (dl *stdLogger) Error(v ...interface{}) {
	dl.output(v...)
}

func (dl *stdLogger) Errorf(fmt string, v ...interface{}) {
	dl.outputf(fmt, v...)
}

func (dl *stdLogger) Fatal(v ...interface{}) {
	dl.output(v...)
}

func (dl *stdLogger) Fatalf(fmt string, v ...interface{}) {
	dl.outputf(fmt, v...)
}

func (dl *stdLogger) Panic(v ...interface{}) {
	dl.output(v...)
}

func (dl *stdLogger) Panicf(fmt string, v ...interface{}) {
	dl.outputf(fmt, v...)
}

func (dl *stdLogger) output(v ...interface{}) {
	os.Stdout.WriteString(dl.packetOutput(fmt.Sprint(v...) + "\n"))
}

func (dl *stdLogger) outputf(format string, v ...interface{}) {
	os.Stdout.WriteString(dl.packetOutput(fmt.Sprintf(format, v...) + "\n"))
}

func (dl *stdLogger) packetOutput(str string) string {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "???"
		line = 0
	}

	f := pathBaseName(runtime.FuncForPC(pc).Name(), 1)

	return "[" + pathBaseName(file, 2) + ":" + strconv.Itoa(line) + "]["+ f+ "] " + str
}

func pathBaseName(path string, deepth int) string {
	d := 1
	for i := len(path) - 1; i > 0; i-- {
		if path[i] == '/' {
			if d == deepth {
				return path[i+1:]
				break
			}
			d++
		}
	}

	return path
}