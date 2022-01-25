package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logStruct struct {
	time    time.Time
	level   string
	message string
}

var logCh = make(chan logStruct, 50)
var doneCh = make(chan struct{})

func main() {
	// Monitor log channel
	go logger()
	defer func() {
		close(logCh)
	}()

	logCh <- logStruct{time.Now(), logInfo, "Starting..."}

	// Do something here
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Do somthing here...")

	logCh <- logStruct{time.Now(), logInfo, "Shutting down..."}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

// Listen to log channel and format it
func logger() {
	for {
		select {
		case log := <-logCh:
			fmt.Printf("%v - [%v]%v\n", log.time.Format("2006-01-02T15:04:05"), log.level, log.message)
		case <-doneCh:
			break
		}
	}
}
