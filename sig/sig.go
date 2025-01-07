package sig

import (
	"os"
	"os/signal"
)

func WaitForSignals(sigs []os.Signal) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, sigs...)
	<-sigChan
}

func WaitUntilSignalled() {
	WaitForSignals([]os.Signal{os.Kill, os.Interrupt})
}
