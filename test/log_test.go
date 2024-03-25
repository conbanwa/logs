package logs_test

import (
	"os"
	"testing"

	"github.com/conbanwa/logs"
)

func TestLog(t *testing.T) {
	os.Setenv("LOG_LEVEL", "WARN")
	logs.I("hello world")
	logs.Blue("hello world")
	logs.Log.Level = logs.L_ERROR
	for i := 0; i < 10; i++ {
		logs.Inline("hello world")
	}
	logs.Green("hello world")
	logs.I("hello world")
	logs.Blue("hello world")
	logs.W("hello world")
}

func TestWrite(t *testing.T) {
	logs.Inline("test")
}
