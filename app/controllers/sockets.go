package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/revel/revel"
	"math/rand"
	"time"
)

type Sockets struct {
	*revel.Controller
}

func (s Sockets) Testing(ws *websocket.Conn) revel.Result {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i += 1 {
		ws.Write([]byte(string(i)))
		time.Sleep(randomTime(time.Millisecond, time.Second))
	}
	return nil
}

func randomTime(min, max time.Duration) time.Duration {
	minInt := int(min)
	maxInt := int(max)
	return time.Duration(rand.Intn(maxInt-minInt) + minInt)
}
