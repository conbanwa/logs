package logs

import (
	"fmt"
	"log"
)

var (
	Log = NewLogger()
	I   = fmt.Sprintf("\033[1;36;40m%s\033[0m", "I")
	D   = fmt.Sprintf("\033[1;32;40m%s\033[0m", "D")
	E   = fmt.Sprintf("\033[1;35;40m%s\033[0m", "E")
	W   = fmt.Sprintf("\033[1;33;40m%s\033[0m", "W")
	P   = fmt.Sprintf("\033[1;35;40m%s\033[0m", "P")
	F   = fmt.Sprintf("\033[1;31;40m%s\033[0m", "F")
	A   = fmt.Sprintf("\033[1;34;40m%s\033[0m", "A")
	C   = fmt.Sprintf("\033[1;34;40m%s\033[0m", "C")
)

const (
	DEBUG Level = iota + 1
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

type Level int

type Logger struct {
	*log.Logger
	level Level
}
