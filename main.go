package logs

import (
	"fmt"
)

// @version 0.3.3
// @license.name last updated at 2023/7/14 23:26:34

func SetLogLevel(level Level) {
	Log.SetLogLevel(level)
}

func Debug(args ...interface{}) {
	Log.output(L_DEBUG, d, Concat(args...))
}

func Debugf(format string, args ...interface{}) {
	Log.output(L_DEBUG, d, fmt.Sprintf(format, args...))
}

func Info(args ...interface{}) {
	Log.output(L_INFO, i, Concat(args...))
}

func Infof(format string, args ...interface{}) {
	Log.output(L_INFO, i, fmt.Sprintf(format, args...))
}

func Warn(args ...interface{}) {
	Log.output(L_WARN, w, Concat(args...))
}

func Warnf(format string, args ...interface{}) {
	Log.output(L_WARN, w, fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	Log.output(L_ERROR, e, Concat(args...))
}

func Errorf(format string, args ...interface{}) {
	Log.output(L_ERROR, e, fmt.Sprintf(format, args...))
}

func Panic(args ...interface{}) {
	panic(Concat(args...))
}

func Panicf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func Fatal(args ...interface{}) {
	panic(Concat(args...))
	exit(1)
}

func Fatalf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
	exit(1)
}

func (l *Logger) SetLogLevel(level Level) {
	l.level = level
}

func (l *Logger) output(le Level, prefix string, log string) {
	if le >= l.level {
		l.Output(3, fmt.Sprintf("%s %s", prefix, log))
	}
	if le >= L_PANIC {
		panic(log)
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.output(L_DEBUG, d, fmt.Sprint(args...))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.output(L_DEBUG, d, fmt.Sprintf(format, args...))
}

func (l *Logger) Info(args ...interface{}) {
	l.output(L_INFO, i, fmt.Sprint(args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.output(L_INFO, i, fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(args ...interface{}) {
	l.output(L_WARN, w, fmt.Sprint(args...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.output(L_WARN, w, fmt.Sprintf(format, args...))
}

func (l *Logger) Error(args ...interface{}) {
	l.output(L_ERROR, e, fmt.Sprint(args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.output(L_ERROR, e, fmt.Sprintf(format, args...))
}

func (l *Logger) Panic(args ...interface{}) {
	panic(Concat(args...))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatal(args ...interface{}) {
	panic(Concat(args...))
	exit(1)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
	exit(1)
}
