package protocol

import (
	"fmt"
	"reflect"
	"testing"
)

// func TestDeserializeMe(t *testing.T) {
// 	testmessages := []string {
// 		"$-1\r\n",
// 		"*1\r\n$4\r\nping\r\n",
// 		"*2\r\n$4\r\necho\r\n$11\r\nhello world\r\n",
// 		"*2\r\n$3\r\nget\r\n$3\r\nkey\r\n",
// 		"+OK\r\n",
// 		"-Error message\r\n",
// 		"$0\r\n\r\n",
// 		"+hello world\r\n",
// 	}
// 	for _, message := range testmessages {
// 		fmt.Println(Deserialize(message))
// 	}
// }

func TestSerializeString(t *testing.T) {
	str := NewString("hello world")
	serializedstring := str.Serialize()
	fmt.Println(serializedstring)
	if serializedstring != "+hello world\r\n" {
		t.Fatalf("%v serialization failed", str)
	}
	t.Fatalf("%s serialization failed", serializedstring)
}

func TestSerializeBulkString(t *testing.T) {
	str := NewBulkString("hello world")
	serializedstring := str.Serialize()
	fmt.Println(serializedstring)
	if serializedstring != "$11\r\nhello world\r\n" {
		t.Fatalf("%s serialization failed", serializedstring)
	}
}

func TestSerializeArray(t *testing.T) {
	
	array := NewArray(
		NewString("hello"),
		NewInteger(5),
		NewError("error happened"),
		NewBulkString("word"),
		//NewArray(NewInteger(51)),
	)
	serializedstring := array.Serialize()
	deserializedArray, err := Deserialize(serializedstring)
	if err != nil || !reflect.DeepEqual(array, deserializedArray) {
		t.Fatalf("%v serialization failed %v", array, deserializedArray)
	}
}