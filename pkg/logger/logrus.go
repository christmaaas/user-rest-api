package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

const logFilePath = "logs/app.log"

var e *logrus.Entry

type logrusLogger struct {
	entry *logrus.Entry
}

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (l *logrusLogger) Debug(args ...any) {
	l.entry.Debug(args...)
}

func (l *logrusLogger) Debugf(format string, args ...any) {
	l.entry.Debugf(format, args...)
}

func (l *logrusLogger) Info(args ...any) {
	l.entry.Info(args...)
}

func (l *logrusLogger) Infof(format string, args ...any) {
	l.entry.Infof(format, args...)
}

func (l *logrusLogger) Warning(args ...any) {
	l.entry.Warning(args...)
}

func (l *logrusLogger) Warningf(format string, args ...any) {
	l.entry.Warningf(format, args...)
}

func (l *logrusLogger) Error(args ...any) {
	l.entry.Error(args...)
}

func (l *logrusLogger) Errorf(format string, args ...any) {
	l.entry.Errorf(format, args...)
}

func (l *logrusLogger) Fatal(args ...any) {
	l.entry.Fatal(args...)
}

func (l *logrusLogger) Fatalf(format string, args ...any) {
	l.entry.Fatalf(format, args...)
}

func (l *logrusLogger) Panic(args ...any) {
	l.entry.Panic(args...)
}

func (l *logrusLogger) Panicf(format string, args ...any) {
	l.entry.Panicf(format, args...)
}

func (l *logrusLogger) Trace(args ...any) {
	l.entry.Trace(args...)
}

func (l *logrusLogger) Tracef(format string, args ...any) {
	l.entry.Tracef(format, args...)
}

func GetLogger() Logger {
	return &logrusLogger{e}
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writer {
		_, err = w.Write([]byte(line))
	}

	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func Init() {
	l := logrus.New()

	l.SetReportCaller(true)

	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	allFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		panic(fmt.Sprintf("Open file error: %s", err))
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
