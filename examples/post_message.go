package example

import (
	"fmt"

	"github.com/magicdrive/go-chatwork"
	rooms "github.com/magicdrive/go-chatwork/api/resource/rooms_sub"
	"github.com/magicdrive/go-chatwork/optional"
)

func PostMessage() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	room_id := 1
	params := rooms.MessagePostParam{Body: "Hi there.", SelfUnread: optional.BoolTrue()}

	if r, err := client.Rooms().Message().Post(room_id, params); err == nil {
		fmt.Println(r.MessageId)
	} else {
		panic(err)
	}

}
