package logs

import (
	"fmt"
	"log"
)

var (
	// Log is the default logger
	Log = NewLogger()
	i   = fmt.Sprintf("\033[1;36;40m%s\033[0m", "I")
	d   = fmt.Sprintf("\033[1;32;40m%s\033[0m", "D")
	e   = fmt.Sprintf("\033[1;35;40m%s\033[0m", "E")
	w   = fmt.Sprintf("\033[1;33;40m%s\033[0m", "W")
	p   = fmt.Sprintf("\033[1;35;40m%s\033[0m", "P")
	f   = fmt.Sprintf("\033[1;31;40m%s\033[0m", "F")
	a   = fmt.Sprintf("\033[1;34;40m%s\033[0m", "a")
	c   = fmt.Sprintf("\033[1;34;40m%s\033[0m", "C")
)

const (
	L_DEBUG Level = iota + 1
	L_INFO
	L_WARN
	L_ERROR
	L_PANIC
	L_FATAL
)

type Level int

type Logger struct {
	*log.Logger
	Level Level
}
