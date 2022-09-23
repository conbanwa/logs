# about

a leveled, colored, and comparable logger for go

# usage

```bash
go get github.com/conbanwa/logs
```

```go
package main

import (
    "os"
    "github.com/conbanwa/logs"
)

func main() {
	os.Setenv("LOG_LEVEL", "DEBUG")
    logs.Info("hello world")
    logs.Blue("hello world")
    logs.Green("hello world")
    logs.NotSame(2+3, '5', "2+3 is not 5")
}
```
