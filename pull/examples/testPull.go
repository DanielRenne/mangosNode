package main

import (
	"github.com/DanielRenne/mangosNode/pull"
	"log"
)

const url = "tcp://127.0.0.1:600"

func main() {
	var node pull.Node

	err := node.Pull(url, handlePushMessage)
	if err != nil {
		log.Printf("Error:  %v", err.Error)
	}

	//Code a forever loop to stop main from exiting.
	for {

	}

}

func handlePushMessage(msg []byte) {
	log.Println(string(msg))
}
