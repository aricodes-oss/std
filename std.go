package std

import (
	"github.com/withmandala/go-log"
	"os"
	"os/signal"
	"syscall"
)

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stderr).WithColor()
}

func WaitForKill() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc
}
