package logs

import (
	"fmt"
)

// @version 0.2.9
// @license.name last updated at 2023/3/11 20:59:41

func SetLogLevel(level Level) {
	Log.SetLogLevel(level)
}

func Debug(args ...interface{}) {
	Log.output(DEBUG, D, Concat(args...))
}

func Debugf(format string, args ...interface{}) {
	Log.output(DEBUG, D, fmt.Sprintf(format, args...))
}

func Info(args ...interface{}) {
	Log.output(INFO, I, Concat(args...))
}

func Infof(format string, args ...interface{}) {
	Log.output(INFO, I, fmt.Sprintf(format, args...))
}

func Warn(args ...interface{}) {
	Log.output(WARN, W, Concat(args...))
}

func Warnf(format string, args ...interface{}) {
	Log.output(WARN, W, fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	Log.output(ERROR, E, Concat(args...))
}

func Errorf(format string, args ...interface{}) {
	Log.output(ERROR, E, fmt.Sprintf(format, args...))
}

func Fatal(args ...interface{}) {
	if Log.level <= FATAL {
		Log.output(FATAL, F, Concat(args...))
		exit(1)
	}
}

func Fatalf(format string, args ...interface{}) {
	if Log.level <= FATAL {
		Log.output(FATAL, F, fmt.Sprintf(format, args...))
		exit(1)
	}
}

func Panic(args ...interface{}) {
	if Log.level <= PANIC {
		Log.output(PANIC, P, Concat(args...))
	}
}

func Panicf(format string, args ...interface{}) {
	if Log.level <= PANIC {
		Log.output(PANIC, P, fmt.Sprintf(format, args...))
	}
}

func (l *Logger) SetLogLevel(level Level) {
	l.level = level
}

func (l *Logger) output(le Level, prefix string, log string) {
	if le >= l.level {
		l.Output(3, fmt.Sprintf("%s %s", prefix, log))
	}
	if le >= PANIC {
		panic(log)
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.output(DEBUG, D, fmt.Sprint(args...))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.output(DEBUG, D, fmt.Sprintf(format, args...))
}

func (l *Logger) Info(args ...interface{}) {
	l.output(INFO, I, fmt.Sprint(args...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.output(INFO, I, fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(args ...interface{}) {
	l.output(WARN, W, fmt.Sprint(args...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.output(WARN, W, fmt.Sprintf(format, args...))
}

func (l *Logger) Error(args ...interface{}) {
	l.output(ERROR, E, fmt.Sprint(args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.output(ERROR, E, fmt.Sprintf(format, args...))
}

func (l *Logger) Fatal(args ...interface{}) {
	if l.level <= FATAL {
		l.output(FATAL, F, fmt.Sprint(args...))
		exit(1)
	}
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	if l.level <= FATAL {
		l.output(FATAL, F, fmt.Sprintf(format, args...))
		exit(1)
	}
}

func (l *Logger) Panic(args ...interface{}) {
	if l.level <= PANIC {
		s := fmt.Sprint(args...)
		l.output(PANIC, P, s)
	}
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.level <= PANIC {
		s := fmt.Sprintf(format, args...)
		l.output(PANIC, P, s)
	}
}
