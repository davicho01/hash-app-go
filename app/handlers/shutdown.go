package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"
)

type Shutdown struct {
	logger *log.Logger
}

func NewShutdown() *Shutdown {
	logger := log.New(os.Stdout, "ShutDown Handler: ", log.LstdFlags)
	return &Shutdown{logger}
}

func (handler *Shutdown) Shutdown(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprint(writer, "System Shutting Down!...")

	syscall.Kill(syscall.Getpid(), syscall.SIGINT)

}
