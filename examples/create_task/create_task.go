package examples

import (
	"fmt"
	"time"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

func main() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	taskAssignAccountID1 := 1
	taskAssignAccountID2 := 2

	date := "2022/01/01"
	t, _ := time.Parse("2006/01/02", date)

	roomID := 1
	params := chatwork.RoomTaskPostParam{
		Body:      "Do it now!",
		Limit:     optional.Int64(t.Unix()),
		LimitType: chatwork.RoomTaskLimitTypeDate,
		ToIDs:     param.IntArray(taskAssignAccountID1, taskAssignAccountID2),
	}

	if r, err := client.Rooms().Tasks().Create(roomID, params); err == nil {
		fmt.Println(r.TaskID)
	} else {
		panic(err)
	}

}
