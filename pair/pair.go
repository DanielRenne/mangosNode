//Package pair implements the pair pattern push node for mangos / nanomsg.
package pair

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pair"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

//Node Structure used to send messages to other nodes.
type Node struct {
	url  string
	sock mangos.Socket
}

type ResponseHandler func(*Node, []byte)

//Start a Listen Pair Node.
func (self *Node) Listen(url string, handler ResponseHandler) error {

	self.url = url

	var err error

	if self.sock, err = pair.NewSocket(); err != nil {
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

//Start a Connection Pair Node.
func (self *Node) Connect(url string, handler ResponseHandler) error {

	self.url = url

	var err error

	if self.sock, err = pair.NewSocket(); err != nil {
		return err
	}
	self.sock.AddTransport(ipc.NewTransport())
	self.sock.AddTransport(tcp.NewTransport())
	if err = self.sock.Dial(url); err != nil {
		return err
	}

	go self.processData(handler)

	return nil
}

//Send a message to a pair node
func (self *Node) Send(payload []byte) error {
	return self.sock.Send(payload)
}

//Handles the pair responses.
func (self *Node) processData(handler ResponseHandler) {

	var msg []byte
	var err error

	for {
		if msg, err = self.sock.Recv(); err != nil {
			continue
		}
		go handler(self, msg)
	}
}
