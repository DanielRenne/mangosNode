# mangosNode rep

Example Code to start a reply node, and request node to send request messages.
###Reply Node
	
	package main
	
	import (
		"github.com/DanielRenne/mangosNode/rep"
		"log"
	)
	
	const url = "tcp://127.0.0.1:600"
	
	func main() {
		var node rep.Node
	
		err := node.Listen(url, 2, handleRequests)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		//Code a forever loop to stop main from exiting.
		for {
	
		}
	
	}
	
	func handleRequests(node *rep.Node, msg []byte) {
		log.Println(string(msg))
		node.Reply([]byte("Replying to Request."))
	}


	
###Request Node

	package main
	
	import (
		"github.com/DanielRenne/mangosNode/req"
		"log"
		"time"
	)
	
	const url = "tcp://127.0.0.1:600"
	
	func main() {
		var node req.Node
	
		err := node.Connect(url, handleReply)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		//Code a forever loop to stop main from exiting.
		for {
			time.Sleep(3 * time.Second)
			go node.Request([]byte("Sending Request"))
		}
	
	}
	
	func handleReply(node *req.Node, msg []byte) {
		log.Println(string(msg))
	}


	