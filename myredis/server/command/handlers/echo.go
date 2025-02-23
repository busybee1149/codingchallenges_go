package handlers

import (
	"fmt"
	"learngo/myredis/server/protocol"
)

const ECHO = "ECHO"

type EchoHandler struct {}

func (p EchoHandler) Execute(value protocol.Array) string{
	elements := value.Elements
	if len(elements) != 2 {
		return protocol.NewError(fmt.Sprintf("Unexpected number of arguments for %s",  ECHO)).Serialize()
	}
	return elements[1].(protocol.BulkString).Serialize()
}