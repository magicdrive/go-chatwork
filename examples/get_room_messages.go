package example

import (
	"fmt"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

func GetRoomMessage() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	room_id := 1
	force_flang := optional.BoolTrue()

	mlist, err := client.Rooms().Message().List(room_id, force_flang)
	if err != nil {
		panic(err)
	}

	for _, item := range mlist {
		fmt.Printf(`Account: "%v"\nBody: "%v"\n\n}`, item.Account.Name, item.Body)
	}

}
