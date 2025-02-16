package server

import (
	"fmt"
	"learngo/myredis/server/protocol"
)


type CommandHandler interface {
	Execute(protocol.Array) string
}

var (
	ECHO_STRING = "ECHO"
	PING = protocol.NewBulkString("PING")
	ECHO = protocol.NewBulkString(ECHO_STRING)
)

type PingHandler struct {}

func (p PingHandler) Execute(value protocol.Array) string{
	return protocol.NewString("PONG").Serialize()
}

type EchoHandler struct {}

func (p EchoHandler) Execute(value protocol.Array) string{
	elements := value.Elements
	if len(elements) != 2 {
		return protocol.NewError(fmt.Sprintf("Unexpected number of arguments for %s",  ECHO_STRING)).Serialize()
	}
	return elements[1].(protocol.BulkString).Serialize()
}

var handlers = map[protocol.BulkString]CommandHandler {
	PING: PingHandler{},
	ECHO: EchoHandler{},
}

func extractCommandString(deserializedValue protocol.Array) protocol.BulkString {
	firstElement := deserializedValue.Elements[0] //first element is command 
	return firstElement.(protocol.BulkString)
}

func Serve(request string) string {
	fmt.Println("Raw payload received", request)
	deserialized, _ := protocol.Deserialize(request)
	fmt.Println("Request payload received", deserialized)
	if commandHandler, ok := handlers[extractCommandString(deserialized)]; ok {
		return commandHandler.Execute(deserialized)
	} else {
		return protocol.NewError("Unable to recognize input").Serialize()
	}
}