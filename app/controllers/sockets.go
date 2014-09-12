package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/revel/revel"
	"math/rand"
	"strconv"
	"time"
)

type Sockets struct {
	*revel.Controller
}

func (s Sockets) Testing(ws *websocket.Conn) revel.Result {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i += 1 {
		istr := strconv.Itoa(i)
		sleepTime := randomTime(time.Millisecond, time.Second)
		revel.INFO.Printf("Sending message to client: \"%s\"", istr)
		revel.INFO.Printf("Sleeping for %s sec", sleepTime)
		ws.Write([]byte(istr))
		time.Sleep(sleepTime)
	}
	return nil
}

func randomTime(min, max time.Duration) time.Duration {
	minInt := int(min)
	maxInt := int(max) / 4
	return time.Duration(rand.Intn(maxInt-minInt) + minInt)
}
