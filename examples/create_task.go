package examples

import (
	"fmt"
	"time"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

func CreateTask() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	task_assign_account_id_1 := 1
	task_assign_account_id_2 := 2

	date := "2022/01/01"
	t, _ := time.Parse("2006/01/02", date)

	room_id := 1
	params := chatwork.RoomTaskPostParam{
		Body:      "Do it now!",
		Limit:     optional.Int64(t.Unix()),
		LimitType: chatwork.RoomTaskLimitTypeDate,
		ToIds:     param.IntArray(task_assign_account_id_1, task_assign_account_id_2),
	}

	if r, err := client.Rooms().Tasks().Create(room_id, params); err == nil {
		fmt.Println(r.TaskId)
	} else {
		panic(err)
	}

}
