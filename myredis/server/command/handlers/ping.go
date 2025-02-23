package handlers

import "learngo/myredis/server/protocol"

type PingHandler struct {}

func (p PingHandler) Execute(value protocol.Array) string{
	return protocol.NewString("PONG").Serialize()
}