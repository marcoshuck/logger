package main

import (
	"fmt"
	"github.com/marcoshuck/logger"
	"os"
)

func main() {
	l := logger.NewLogger(logger.VerbosityDebug, os.Stdout)
	msg := "Hello world!"
	l.Debug(fmt.Sprintf("Important message: %s", msg))
}
