//Package push implements the pipeline pattern push node for mangos / nanomsg.
package push

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/push"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

//
type Node struct {
	url  string
	sock mangos.Socket
}
