package main

import (
	"fmt"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api"
	rooms "github.com/magicdrive/go-chatwork/api/resource/rooms_sub"
)

func main() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	room_id := 1
	params := rooms.MessagePostParam{Body: "Hi there.", SelfUnread: api.ChatworkBoolTrue}

	if r, err := client.Rooms().Message().Post(room_id, params); err == nil {
		fmt.Println(r.MessageId)
	} else {
		panic(err)
	}

}
