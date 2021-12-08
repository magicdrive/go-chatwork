package examples

import (
	"fmt"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

func main() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	roomID := 1
	force := optional.BoolTrue()

	mlist, err := client.Rooms().Messages().List(roomID, force)
	if err != nil {
		panic(err)
	}

	for _, item := range mlist {
		fmt.Printf("Account: '%v'\nBody: '%v'\n\n}", item.Account.Name, item.Body)
	}

}
