//Package push implements the pipeline pattern push node for mangos / nanomsg.
package pull

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pull"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

//Node Structure used to send messages to other nodes.
type Node struct {
	url  string
	sock mangos.Socket
}

type ResponseHandler func([]byte)

//Connect to a pull Server.
func (self *Node) Pull(url string, handler ResponseHandler) error {

	self.url = url

	var err error

	if self.sock, err = pull.NewSocket(); err != nil {
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

//Handles the push responses.
func (self *Node) processData(handler ResponseHandler) {

	var msg []byte
	var err error

	for {
		if msg, err = self.sock.Recv(); err != nil {
			continue
		}
		go handler(msg)
	}
}
