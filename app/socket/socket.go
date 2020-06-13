package socket

import (
	socketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/koushamad/goro/app/controller"
	"sync"
)

type (
	Socket struct{ *socketio.Server }
)

var (
	socket *Socket
	once   sync.Once
)

func Init() *Socket {
	once.Do(func() {
		socket = &Socket{socketio.NewServer(transport.GetDefaultWebsocketTransport())}
	})
	return socket
}

func (so *Socket) Handler() {

	so.On(socketio.OnConnection, controller.Connect)

	so.On(socketio.OnDisconnection, controller.Disconnet)

	so.On(socketio.OnError, controller.Error)

	so.On("auth", controller.Auth)

	so.On("join", controller.Join)

	so.On("leave", controller.Leave)

	so.On("send", controller.Send)
}
