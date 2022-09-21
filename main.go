package logs

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// @version 0.0.2
// @description last updated at 9/21/2022 11:57:00 PM
type Level int

type Logger struct {
	*log.Logger
	level Level
}

var (
	Log = NewLogger()
	I   = "I"
	D   = fmt.Sprintf("\033[0;32;40m%s\033[0m", "D")
	E   = fmt.Sprintf("\033[0;35;40m%s\033[0m", "E")
	W   = fmt.Sprintf("\033[1;33;40m%s\033[0m", "W")
	F   = fmt.Sprintf("\033[1;31;40m%s\033[0m", "F")
	P   = fmt.Sprintf("\033[1;31;40m%s\033[0m", "P")
	A   = fmt.Sprintf("\033[1;31;40m%s\033[0m", "A")
	C   = fmt.Sprintf("\033[1;31;40m%s\033[0m", "C")
)

const (
	DEBUG Level = iota + 1
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

func NewLogger() *Logger {
	if os.Getenv("PORT") == "8080" {
		return &Logger{
			Logger: log.New(os.Stderr, "", log.Lshortfile|log.Ltime|log.Lmsgprefix|log.Lmicroseconds),
			level:  DEBUG,
		}
	}
	return &Logger{
		Logger: log.New(os.Stderr, "", log.Llongfile|log.Ltime|log.Lmsgprefix),
		level:  DEBUG,
	}
}
func init() {
	logLevel := os.Getenv("LOG_LEVEL")
	logFileName := os.Getenv("LOG_FILE")
	var l Level
	switch strings.ToLower(logLevel) {
	case "debug", "DEBUG":
		l = DEBUG
	case "info", "INFO":
		l = INFO
	case "warn", "WARN":
		l = WARN
	case "error", "ERROR":
		l = ERROR
	case "fatal", "FATAL":
		l = FATAL
	case "panic", "PANIC":
		l = PANIC
	default:
		l = DEBUG
	}
	SetLevel(l)
	if logFileName != "" {
		f, err := os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
		if err == nil {
			SetOut(f)
		} else {
			Warn("log file not open ??? ")
			Error(err.Error())
		}
	}
}

func Assert(b bool, args ...interface{}) {
	if !b {
		Log.output(PANIC, A, Concat(args...))
		os.Exit(1)
	}
}
func MustEqual(a, b interface{}, args ...interface{}) {
	if a != b {
		Log.output(PANIC, A, Concat(append(args, a, "not equal to", b)))
		os.Exit(1)
	}
}
func ExitIfErr(err error, args ...interface{}) {
	if err != nil {
		Log.output(PANIC, C, err.Error()+Concat(args...))
		os.Exit(1)
	}
}
func NotNil(err error, args ...interface{}) bool {
	if err != nil {
		Log.output(ERROR, E, fmt.Sprint(err, Concat(args...)))
		return true
	}
	return false
}

func Concat(args ...interface{}) (str string) {
	return ConcatWith(" ", args...)
}
func ConcatWith(separator string, args ...interface{}) (str string) {
	for i, param := range args {
		if i == len(args)-1 {
			separator = ""
		}
		str += fmt.Sprint(param) + separator
	}
	return
}
func Dye(highlight int, color string, args ...interface{}) string {
	str := Concat(args...)
	n := "37"
	switch color {
	case "red":
		n = "31"
	case "green":
		n = "32"
	case "yellow":
		n = "33"
	case "blue":
		n = "34"
	case "magenta":
		n = "35"
	case "cyan":
		n = "36"
	default:
		n = "33"
	}
	return fmt.Sprintf("\033["+strconv.Itoa(highlight)+";"+n+";40m%s\033[0m", str)
}
func Highlight(color string, args ...interface{}) {
	Log.output(INFO, I, Dye(1, color, args...))
}
func Red(args ...interface{}) {
	Log.output(INFO, I, Dye(0, "red", args...))
}
func Green(args ...interface{}) {
	Log.output(INFO, I, Dye(0, "green", args...))
}
func Yellow(args ...interface{}) {
	Log.output(INFO, I, Dye(0, "yellow", args...))
}
func Blue(args ...interface{}) {
	Log.output(INFO, I, Dye(0, "blue", args...))
}
func Magenta(args ...interface{}) {
	Log.output(INFO, I, Dye(0, "magenta", args...))
}
func Cyan(args ...interface{}) {
	Log.output(INFO, I, Dye(0, "cyan", args...))
}
func SetOut(out io.Writer) {
	Log.SetOut(out)
}
func SetLevel(level Level) {
	Log.SetLevel(level)
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
		os.Exit(1)
	}
}
func Fatalf(format string, args ...interface{}) {
	if Log.level <= FATAL {
		Log.output(FATAL, F, fmt.Sprintf(format, args...))
		os.Exit(1)
	}
}
func Panic(args ...interface{}) {
	if Log.level <= PANIC {
		Log.output(PANIC, P, Concat(args...))
		panic(Concat(args...))
	}
}
func Panicf(format string, args ...interface{}) {
	if Log.level <= PANIC {
		Log.output(PANIC, P, fmt.Sprintf(format, args...))
		panic(Concat(args...))
	}
}
func (l *Logger) SetLevel(level Level) {
	l.level = level
}
func (l *Logger) SetOut(out io.Writer) {
	l.Logger.SetOutput(out)
}
func (l *Logger) output(le Level, prefix string, log string) {
	if l.level <= le {
		l.Output(3, fmt.Sprintf("%s %s", prefix, log))
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
		os.Exit(1)
	}
}
func (l *Logger) Fatalf(format string, args ...interface{}) {
	if l.level <= FATAL {
		l.output(FATAL, F, fmt.Sprintf(format, args...))
		os.Exit(1)
	}
}
func (l *Logger) Panic(args ...interface{}) {
	if l.level <= PANIC {
		s := fmt.Sprint(args...)
		l.output(PANIC, P, s)
		panic(s)
	}
}
func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.level <= PANIC {
		s := fmt.Sprintf(format, args...)
		l.output(PANIC, P, s)
		panic(s)
	}
}