package main

import (
	".../server"
	//"./Network/localip"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	enableCh := make(chan bool)
	sendCh := make(chan server.Message)
	receiveCh := make(chan server.Message)
	eventCh := make(chan server.Event)


	var id string
	id = "test";

	go server.Server(20013, id, enableCh, sendCh, receiveCh, eventCh)

	
	enableCh <- true
	
	for {
		select {
		case message := <-receiveCh:
			fmt.Println("Received: ", message)
		case event := <-eventCh:
			fmt.Println("Event: ", server.ToString(event))
		default:
			//sendCh <- "hei"
			time.Sleep(50*time.Millisecond)
		}
	}
}
