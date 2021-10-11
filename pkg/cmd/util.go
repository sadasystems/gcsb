package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
)

// graceful wraps a context cancel func with a listener for OS interrupt signals
func graceful(cancelFn context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		oscall := <-c
		log.Printf("System call received. Exiting! (%s)", oscall)
		cancelFn()
	}()
}
