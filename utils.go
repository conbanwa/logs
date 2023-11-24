package logs

import (
	"fmt"
	"net"
	"io"
	"log"
	"os"
	"strconv"
)

func init() {
	logLevel := os.Getenv("LOG_LEVEL")
	logFileName := os.Getenv("LOG_FILE")
	SetLevel(logLevel)
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

func NewLogger() *Logger {
	if os.Getenv("ENV") != "" {
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

func SetOut(out io.Writer) {
	Log.SetOut(out)
}

func (l *Logger) SetOut(out io.Writer) {
	l.Logger.SetOutput(out)
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

func Err(args ...interface{}) error {
	str := Concat(args...)
	return fmt.Errorf(str)
}

// SetLogLevel by string
func SetLevel(level string) {
	SetLogLevel(StringLevel(level))
}

func StringLevel(level string) Level {
	switch level {
	case "debug", "DEBUG":
		return DEBUG
	case "info", "INFO":
		return INFO
	case "warn", "WARN":
		return WARN
	case "error", "ERROR":
		return ERROR
	case "fatal", "FATAL":
		return FATAL
	case "panic", "PANIC":
		return PANIC
	default:
		return DEBUG
	}
}

func ParseFloat64(a interface{}) (f float64, err error) {
	switch a := a.(type) {
	case int:
		f = float64(a)
	case int8:
		f = float64(a)
	case int16:
		f = float64(a)
	case int32:
		f = float64(a)
	case int64:
		f = float64(a)
	case uint:
		f = float64(a)
	case uint8:
		f = float64(a)
	case uint16:
		f = float64(a)
	case uint32:
		f = float64(a)
	case uint64:
		f = float64(a)
	case float32:
		f = float64(a)
	case float64:
		f = a
	case string:
		f, err = strconv.ParseFloat(a, 64)
	default:
		err = fmt.Errorf("not a number")
	}
	return
}

func ToFloat64(a interface{}) float64 {
	f, err := ParseFloat64(a)
	if err != nil {
		Fatal(err)
		return 0
	}
	return f
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

func Sprint(args ...interface{}) string {
	Log.output(INFO, I, Concat(args...))
	return Concat(args...)
}

func exit(code int) {
	os.Exit(code)
}

func Inline(args ...interface{}) {
	str := Concat(args...)
	//output to stdout
	b := []byte(str)
	os.Stderr.Write(b)
}

func IpList() []string {
	var ips []string
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		ErrorIfNotNil(err)
	}
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

func Table[T any](arr []T, args ...any) {
	out := fmt.Sprint(args...) + "\n"
	for i, a := range arr {
		out += fmt.Sprintf("%d|%+v\n", i, a)
	}
	Log.output(INFO, I, out)
}
