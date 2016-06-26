package main

import (
	"github.com/DanielRenne/mangosNode/pair"
	"log"
	"time"
)

const url = "tcp://127.0.0.1:600"

func main() {
	var node pair.Node

	err := node.Listen(url, handlePairMessage)
	if err != nil {
		log.Printf("Error:  %v", err.Error)
	}

	//Code a forever loop to stop main from exiting.
	for {
		time.Sleep(3 * time.Second)
		go node.Send([]byte("Sending Data from Listener"))
	}

}

func handlePairMessage(msg []byte) {
	log.Println(string(msg))
}
