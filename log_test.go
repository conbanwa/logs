package logs_test

import (
	"os"
	"testing"

	"github.com/conbanwa/logs"
)

func TestLog(t *testing.T) {
	os.Setenv("LOG_LEVEL", "DEBUG")
	logs.Info("hello world")
	logs.Blue("hello world")
	for i := 0; i < 10; i++ {
		logs.Write("hello world")
	}
	logs.Green("hello world")
	logs.NotSame(2+3, '5', "2+3 is not 5")
}
