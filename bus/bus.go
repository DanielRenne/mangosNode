//Package bus implements the bus pattern push node for mangos / nanomsg.
package bus

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/bus"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"log"
)

//Node Structure used to send messages to other nodes.
type Node struct {
	url  string
	sock mangos.Socket
}

type ResponseHandler func([]byte)

//Start a Listen Bus Node.
func (self *Node) Listen(url string, handler ResponseHandler) error {

	self.url = url

	var err error

	if self.sock, err = bus.NewSocket(); err != nil {
		return err
	}
	self.sock.AddTransport(ipc.NewTransport())
	self.sock.AddTransport(tcp.NewTransport())
	if err = self.sock.Listen(url); err != nil {
		return err
	}

	go self.processData(handler)

	return nil
}

//Start a Connection to another bus Node.
func (self *Node) Connect(url string) error {

	var err error

	if err = self.sock.Dial(url); err != nil {
		return err
	}

	return nil
}

//Send a message to a bus node
func (self *Node) Send(payload []byte) error {
	return self.sock.Send(payload)
}

//Handles the bus responses.
func (self *Node) processData(handler ResponseHandler) {

	var msg []byte
	var err error

	for {
		if msg, err = self.sock.Recv(); err != nil {
			continue
		}
		log.Println(string(msg))
		go handler(msg)
	}
}
