package main

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/examples/chat/types"
	"github.com/anthdm/hollywood/log"
	"github.com/anthdm/hollywood/remote"
)

type server struct {
	clients map[*actor.PID]string
}

func newServer() actor.Receiver {
	return &server{
		clients: make(map[*actor.PID]string),
	}
}
func (s *server) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case *types.Message:
		s.handleMessage(ctx, msg)
	case *types.Connect:
		s.clients[ctx.Sender()] = msg.Username
		log.Infow("Новый клиет подсоединился", log.M{
			"pid":      ctx.Sender(),
			"username": msg.Username,
		})
	case actor.Started:
	case actor.Stopped:
		_ = msg

	}
}

func (s *server) handleMessage(ctx *actor.Context, msg *types.Message) {
	for pid := range s.clients {
		if !pid.Equals(ctx.Sender()) {
			ctx.Forward(pid)
		}
	}
}

func main() {
	log.SetLevel(log.LevelInfo)
	e := actor.NewEngine()
	rem := remote.New(e, remote.Config{
		ListenAddr: "127.0.0.1:4000",
	})
	e.WithRemote(rem)
	e.Spawn(newServer, "server")
	<-(make(chan struct{}))
}
