package exithandler

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Init(cb func()){
	sigs := make(chan os.Signal, 1)
	terminate := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println("exit reason: ", sig)
		close(terminate)
	}()
	
	<-terminate
	cb()
}