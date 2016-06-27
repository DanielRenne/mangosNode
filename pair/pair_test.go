package pair

import (
	"log"
	"testing"
)

const url = "tcp://127.0.0.1:600"

var tGlobal *testing.T
var messages chan string
var messages2 chan string

func TestPair(t *testing.T) {
	tGlobal = t
	messages = make(chan string)
	messages2 = make(chan string)

	var nodePairListener Node
	var nodePairConnection Node

	err := nodePairListener.Listen(url, handlePairResponse)

	if err != nil {
		t.Errorf("Error starting listen pair node at pair_test.TestPair:  %v", err.Error())
		return
	}

	err = nodePairConnection.Connect(url, handlePairResponseConnection)

	if err != nil {
		t.Errorf("Error connecting connection node at pair_test.TestPair:  %v", err.Error())
		return
	}

	err = nodePairConnection.Send([]byte("TestingPair1"))

	if err != nil {
		t.Errorf("Error sending message at pair_test.TestPair:  %v", err.Error())
		return
	}

	msg := <-messages
	log.Println(msg)

	msg2 := <-messages2
	log.Println(msg2)
}

func handlePairResponse(node *Node, msg []byte) {

	if string(msg) != "TestingPair1" {
		tGlobal.Errorf("Failed to match the push response message at pair_test.handlePairResponse")
		messages <- "Test Client 1 Failed"
		return
	}

	messages <- "Test Client 1 Passed"

	err := node.Send([]byte("TestingPair2"))

	if err != nil {
		tGlobal.Errorf("Error sending message at pair_test.handlePairRepsonse:  %v", err.Error())
		return
	}
}

func handlePairResponseConnection(node *Node, msg []byte) {

	if string(msg) != "TestingPair2" {
		tGlobal.Errorf("Failed to match the push response message at pair_test.handlePairResponseConnection")
		messages2 <- "Test Client 2 Failed"
		return
	}

	messages2 <- "Test Client 2 Passed"
}
