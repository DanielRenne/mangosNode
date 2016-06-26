package push

import (
	"github.com/DanielRenne/mangosNode/pull"
	"log"
	"testing"
)

const url = "tcp://127.0.0.1:600"

var tGlobal *testing.T
var messages chan string

func TestPush(t *testing.T) {
	tGlobal = t
	messages = make(chan string)

	var nodeP pull.Node
	err := nodeP.Pull(url, handlePullResponse)

	if err != nil {
		t.Errorf("Error starting pull node at push_test.TestPush:  %v", err.Error())
		return
	}

	var nodePush Node
	err = nodePush.Connect(url)

	if err != nil {
		t.Errorf("Error connecting push node at push_test.TestPush:  %v", err.Error())
		return
	}

	err = nodePush.Push([]byte("TestingPush"))

	if err != nil {
		t.Errorf("Error pushing message at push_test.TestPush:  %v", err.Error())
		return
	}

	msg := <-messages
	log.Println(msg)
}

func handlePullResponse(msg []byte) {

	log.Println(string(msg))

	if string(msg) != "TestingPush" {
		tGlobal.Errorf("Failed to match the push response message at push_test.handlePullResponse")
		messages <- "Test Failed"
		return
	}

	messages <- "Test Passed"

}
