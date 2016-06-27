//Package req provides implementation of a request mangos node.
package req

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/req"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

type Node struct {
	url  string
	sock mangos.Socket
}

type ResponseHandler func(*Node, []byte)

//Starts a request Node on the specified url.
func (self *Node) Connect(url string, handler ResponseHandler) error {

	self.url = url

	var err error

	if self.sock, err = req.NewSocket(); err != nil {
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

//Request the reply node.
func (self *Node) Request(payload []byte) error {

	var err error

	if err = self.sock.Send(payload); err != nil {
		return err
	}

	return nil
}

//Handles the reply Messages.
func (self *Node) processData(handler ResponseHandler) {

	for {

		msg, err := self.sock.Recv()
		if err != nil {
			continue
		}
		go handler(self, msg)
	}
}
