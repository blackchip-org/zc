package repl

import (
	"os"
	"os/signal"
	"syscall"
)

func consoleInit(r *REPL) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		r.Close()
	}()
}
