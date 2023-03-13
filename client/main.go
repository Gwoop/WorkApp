package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/examples/chat/types"
	"github.com/anthdm/hollywood/log"
	"github.com/anthdm/hollywood/remote"
	"os"
)

type client struct {
	username  string
	serverPID *actor.PID
}

func newClient(username string, serverPID *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &client{
			username:  username,
			serverPID: serverPID,
		}
	}
}
func (c *client) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case *types.Message:
		fmt.Println("username: %s :: %s\n", msg.Username, msg.Msg)
	case actor.Started:
		ctx.Send(c.serverPID, &types.Connect{
			Username: c.username,
		})
	case actor.Stopped:
		_ = msg
	}
}

func main() {
	var (
		port     = flag.String("port", ":3000", "")
		username = flag.String("username", "", "")
	)
	flag.Parse()
	e := actor.NewEngine()
	rem := remote.New(e, remote.Config{
		ListenAddr: "127.0.0.1" + *port,
	})
	e.WithRemote(rem)
	serverPID := actor.NewPID("127.0.0.1:4000", "server")
	clientPID := e.Spawn(newClient(*username, serverPID), "client")
	r := bufio.NewReader(os.Stdin)
	for {
		str, err := r.ReadString('\n')
		if err != nil {
			log.Errorw("ошибка чтения сообщения от stdin (cl/main/47)", log.M{"err": err})
		}
		msg := &types.Message{
			Msg:      str,
			Username: *username,
		}
		e.SendWithSender(serverPID, msg, clientPID)
	}
}
