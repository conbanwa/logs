package logs

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
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
			W("log file not open ??? ")
			E(err.Error())
		}
	}
}

func NewLogger() *Logger {
	if os.Getenv("ENV") != "" {
		return &Logger{
			Logger: log.New(os.Stderr, "", log.Lshortfile|log.Ltime|log.Lmsgprefix|log.Lmicroseconds),
			Level:  L_DEBUG,
		}
	}
	return &Logger{
		Logger: log.New(os.Stderr, "", log.Llongfile|log.Ltime|log.Lmsgprefix),
		Level:  L_DEBUG,
	}
}

func SetOut(out io.Writer) {
	Log.SetOut(out)
}

func (l *Logger) SetOut(out io.Writer) {
	l.Logger.SetOutput(out)
}

func Concat(args ...interface{}) string {
	return ReplaceBool(ConcatWith(" ", args...))
}

func ReplaceBool(old string) (str string) {
	str = strings.Replace(old, "true", "√", -1)
	str = strings.Replace(str, "false", "x", -1)
	return
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
	Log.Level = StringLevel(level)
}

func StringLevel(level string) Level {
	switch level {
	case "debug", "DEBUG":
		return L_DEBUG
	case "info", "INFO":
		return L_INFO
	case "warn", "WARN":
		return L_WARN
	case "error", "ERROR":
		return L_ERROR
	case "fatal", "FATAL":
		return L_FATAL
	case "panic", "PANIC":
		return L_PANIC
	default:
		return L_DEBUG
	}
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
	Log.output(L_INFO, i, Concat(args...))
	return Concat(args...)
}

func exit(code int) {
	os.Exit(code)
}

func Inline(args ...interface{}) {
	if Log.Level > L_DEBUG {
		return
	}
	str := Concat(args...) + "|"
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
