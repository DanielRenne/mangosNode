package bus

import (
	"log"
	"testing"
	"time"
)

const url1 = "tcp://127.0.0.1:600"
const url2 = "tcp://127.0.0.1:601"
const url3 = "tcp://127.0.0.1:602"

var tGlobal *testing.T
var messages1 chan string
var messages2 chan string
var messages3 chan string

func TestBus(t *testing.T) {
	tGlobal = t
	messages1 = make(chan string)
	messages2 = make(chan string)
	messages3 = make(chan string)

	var bus1 Node
	var bus2 Node
	var bus3 Node

	err := bus1.Listen(url1, handleBus1Response)

	if err != nil {
		t.Errorf("Error starting bus1 node at bus_test.TestBus:  %v", err.Error())
		return
	}

	err = bus2.Listen(url2, handleBus2Response)

	if err != nil {
		t.Errorf("Error starting bus2 node at bus_test.TestBus:  %v", err.Error())
		return
	}

	err = bus3.Listen(url3, handleBus3Response)

	if err != nil {
		t.Errorf("Error starting bus3 node at bus_test.TestBus:  %v", err.Error())
		return
	}

	err = bus1.Connect(url2)

	if err != nil {
		t.Errorf("Error connecting bus1 node at bus_test.TestBus:  %v", err.Error())
		return
	}

	err = bus2.Connect(url3)

	if err != nil {
		t.Errorf("Error connecting bus2 node at bus_test.TestBus:  %v", err.Error())
		return
	}

	err = bus3.Connect(url1)

	if err != nil {
		t.Errorf("Error connecting bus3 node at bus_test.TestBus:  %v", err.Error())
		return
	}
	time.Sleep(time.Second)

	err = bus1.Send([]byte("BusMessage"))

	if err != nil {
		t.Errorf("Error sending message at bus_test.TestBus:  %v", err.Error())
		return
	}

	msg2 := <-messages2
	log.Println(msg2)

	msg3 := <-messages3
	log.Println(msg3)

	err = bus2.Send([]byte("BusMessage"))

	if err != nil {
		t.Errorf("Error sending message at bus_test.TestBus:  %v", err.Error())
		return
	}

	msg1 := <-messages1
	log.Println(msg1)

}

func handleBus1Response(msg []byte) {

	log.Println(string(msg))
	if string(msg) != "BusMessage" {
		tGlobal.Errorf("Failed to match the Bus response message at bus_test.handleBus1Response")
		messages1 <- "Bus 1 Response Failed"
		return
	}

	messages1 <- "Bus 1 Passed"
}

func handleBus2Response(msg []byte) {

	log.Println(string(msg))
	if string(msg) != "BusMessage" {
		tGlobal.Errorf("Failed to match the Bus response message at bus_test.handleBus2Response")
		messages2 <- "Bus 2 Response Failed"
		return
	}

	messages2 <- "Bus 2 Passed"
}

func handleBus3Response(msg []byte) {

	log.Println(string(msg))
	if string(msg) != "BusMessage" {
		tGlobal.Errorf("Failed to match the Bus response message at bus_test.handleBus3Response")
		messages3 <- "Bus 3 Response Failed"
		return
	}

	messages3 <- "Bus 3 Passed"
}
