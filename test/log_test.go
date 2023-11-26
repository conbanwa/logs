package logs_test

import (
	"os"
	"testing"

	"github.com/conbanwa/logs"
)

func TestLog(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	logs.I("hello world")
	logs.Blue("hello world")
	for i := 0; i < 10; i++ {
		logs.Inline("hello world")
	}
	logs.Green("hello world")
	logs.I("hello world")
	logs.Blue("hello world")
}
