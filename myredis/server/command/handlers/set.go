package handlers

import (
	"fmt"
	"learngo/myredis/server/protocol"
	"learngo/myredis/server/store"
)

const (
	SET = "SET"
     OK = "OK"
)

type SetHandler struct {
	KeyValueStore *store.KeyValueStore
}

func (s SetHandler) Execute(value protocol.Array) string {
	elements := value.Elements
	if len(elements) != 3 {
		return protocol.NewError(fmt.Sprintf("Unexpected number of arguments for %s",  SET)).Serialize()
	}

	setKey, setKeyOk := elements[1].(protocol.BulkString)
	setValue, setValueOk := elements[2].(protocol.BulkString)

	if !setKeyOk || !setValueOk {
		return protocol.NewError(fmt.Sprintf("Unexpected error for %s",  SET)).Serialize()
	}

	//fmt.Printf("%s value %s for key %s", SET_STRING, setKey.Serialize(), setValue.Serialize())
	setStore := s.KeyValueStore
	(*setStore).Set(setKey.ContentString(), setValue.ContentString())
	return protocol.NewString(OK).Serialize()
}