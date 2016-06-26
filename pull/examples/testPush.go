package main

import (
	"github.com/DanielRenne/mangosNode/push"
	"log"
	"time"
)

const url = "tcp://127.0.0.1:600"

func main() {
	var node push.Node

	err := node.Connect(url)
	if err != nil {
		log.Printf("Error:  %v", err.Error)
	}

	//Code a forever loop to stop main from exiting.
	for {
		time.Sleep(3 * time.Second)
		go node.Push([]byte("Pushing Data"))
	}

}
