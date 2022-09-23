# about

a structured, leveled, colored, and formatted logger for go

# usage

```bash
go get github.com/conbanwa/logs
```

```go
package main
import (
    "github.com/conbanwa/logs"
)
func main() {
    logs.Info("hello world")
    logs.Blue("hello world")
    logs.Green("hello world")
    logs.NotSame(2+3, '5', "2+3 is not 5")
}
```
