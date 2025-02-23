package server

import (
	"learngo/myredis/server/protocol"
	"learngo/myredis/server/store"
	"learngo/myredis/server/command/handlers"
)


type CommandHandler interface {
	Execute(protocol.Array) string
}

var (
	GET = protocol.NewBulkString("GET")
	PING = protocol.NewBulkString("PING")
	ECHO = protocol.NewBulkString("ECHO")
	SET = protocol.NewBulkString("SET")
	kvStore = store.NewInMemoryKeyValueStore()
)

var handlerMap = map[protocol.BulkString]CommandHandler {
	PING: handlers.PingHandler{},
	ECHO: handlers.EchoHandler{},
	GET: handlers.GetHandler{ KeyValueStore: &kvStore },
	SET: handlers.SetHandler{ KeyValueStore: &kvStore },
}

func extractCommandString(deserializedValue protocol.Array) protocol.BulkString {
	firstElement := deserializedValue.Elements[0] //first element is command 
	return firstElement.(protocol.BulkString)
}

func Execute(request string) string {
	//fmt.Println("Raw payload received", request)
	deserialized, _ := protocol.Deserialize(request)
	//fmt.Println("Request payload received", deserialized)
	if commandHandler, ok := handlerMap[extractCommandString(deserialized)]; ok {
		return commandHandler.Execute(deserialized)
	} else {
		return protocol.NewError("Unable to recognize input").Serialize()
	}
}