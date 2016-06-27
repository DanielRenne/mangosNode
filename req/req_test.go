package req

import (
	"github.com/DanielRenne/mangosNode/rep"
	"log"
	"testing"
)

const url = "tcp://127.0.0.1:600"

var tGlobal *testing.T
var messages chan string
var messages2 chan string

func TestReq(t *testing.T) {
	tGlobal = t
	messages = make(chan string)
	messages2 = make(chan string)

	var replyNode rep.Node
	var requestNode Node

	err := replyNode.Listen(url, 2, handleRequests)

	if err != nil {
		t.Errorf("Error starting listen reply node at req_test.TestReq:  %v", err.Error())
		return
	}

	err = requestNode.Connect(url, handleReply)

	if err != nil {
		t.Errorf("Error connecting request node at req_test.TestReq:  %v", err.Error())
		return
	}

	err = requestNode.Request([]byte("MyRequest"))

	if err != nil {
		t.Errorf("Error sending request at req_test.TestReq:  %v", err.Error())
		return
	}

	msg := <-messages
	log.Println(msg)

	msg2 := <-messages2
	log.Println(msg2)
}

func handleRequests(node *rep.Node, msg []byte) {

	if string(msg) != "MyRequest" {
		tGlobal.Errorf("Failed to match the reply response message at req_test.handleRequests")
		messages <- "Handle Requests Failed"
		return
	}

	messages <- "Handle Requests Passed"

	err := node.Reply([]byte("MyReply"))

	if err != nil {
		tGlobal.Errorf("Error sending reply message at req_test.handleRequests:  %v", err.Error())
		return
	}

}

func handleReply(node *Node, msg []byte) {

	if string(msg) != "MyReply" {
		tGlobal.Errorf("Failed to match the reply response message at req_test.handleReply")
		messages2 <- "Handle Reply Failed"
		return
	}

	messages2 <- "Handle Reply Passed"
}
