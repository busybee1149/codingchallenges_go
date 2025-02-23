package handlers

import (
	"fmt"
	"learngo/myredis/server/protocol"
	"learngo/myredis/server/store"
)

const GET = "GET"

type GetHandler struct {
	KeyValueStore *store.KeyValueStore
}

func (g GetHandler) Execute(value protocol.Array) string {
	elements := value.Elements
	if len(elements) != 2 {
		return protocol.NewError(fmt.Sprintf("Unexpected number of arguments for %s",  GET)).Serialize()
	}
	if key, ok := elements[1].(protocol.BulkString); ok {
		getstore := g.KeyValueStore
		value, err := (*getstore).Get(key.ContentString())
		if err != nil {
			return protocol.NULL_BULK_STRING.Serialize()
		}
		return protocol.NewString(value).Serialize()
	}
	
	return protocol.NewError(fmt.Sprintf("Unexpected error for %s", GET)).Serialize()
}