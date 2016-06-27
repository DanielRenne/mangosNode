//Package reply provides implementation of a reply mangos node.
package rep

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/rep"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

type Node struct {
	url  string
	sock mangos.Socket
}

type ResponseHandler func(*Node, []byte)

//Starts a Node reply server on the specified url.  A set of workers can run to handle more traffic.
func (self *Node) Listen(url string, workers int, handler ResponseHandler) error {

	self.url = url

	var err error

	if self.sock, err = rep.NewSocket(); err != nil {
		return err
	}

	self.sock.AddTransport(ipc.NewTransport())
	self.sock.AddTransport(tcp.NewTransport())

	if err = self.sock.Listen(url); err != nil {
		return err
	}

	for id := 0; id < workers; id++ {
		go self.processData(handler)
	}

	return nil

}

//Reply to the Request node Message.
func (self *Node) Reply(payload []byte) error {

	var err error

	if err = self.sock.Send(payload); err != nil {
		return err
	}

	return nil
}

//Handles the Request Messages.
func (self *Node) processData(handler ResponseHandler) {

	for {

		msg, err := self.sock.Recv()
		if err != nil {
			continue
		}
		go handler(self, msg)
	}
}
