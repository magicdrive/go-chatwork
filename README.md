# go-chatwork

[![CircleCI](https://circleci.com/gh/magicdrive/go-chatwork/tree/main.svg?style=svg)](https://circleci.com/gh/magicdrive/go-chatwork/tree/main)
[![GoDoc](https://godoc.org/github.com/magicdrive/go-chatwork?status.svg)](https://godoc.org/github.com/magicdrive/go-chatwork)
[![Go Report Card](https://goreportcard.com/badge/github.com/magicdrive/go-chatwork)](https://goreportcard.com/report/github.com/magicdrive/go-chatwork)
[![codecov](https://codecov.io/gh/magicdrive/go-chatwork/branch/main/graph/badge.svg?token=RWP0GEUWPA)](https://codecov.io/gh/magicdrive/go-chatwork)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/magicdrive/go-chatwork/blob/main/LICENSE)

Chatwork API Client for Golang. (API endpoint full supported.)

## Introduction
The Chatwork API Client for Go makes it easy to develop bots using Chatwork API.

## Documentation

See the official API documentation for more information.

* Official API Document:
	- en: http://download.chatwork.com/ChatWork_API_Documentation.pdf
	- ja: http://developer.chatwork.com/ja/
* Github: https://github.com/chatwork/api
* RAML: https://github.com/chatwork/api/blob/master/RAML/api-ja.raml

## Requirements

This library requires Go 1.11 or later.

## Installation ##

```sh
$ go get -u github.com/magicdrive/go-chatwork@latest
```

## Configuration ##

```go
import (
    "github.com/magicdrive/go-chatwork"
)

func main() {
	client := chatwork.NewChatworkClient(`your-api-key-here`)
	...
}

```

## Configuration With Detailed ##

```go
import (
    "net/http"
    "github.com/magicdrive/go-chatwork"
)

func main() {
	http_client := &http.Client{}
	alt_host := "audit-chatwork-api.yourcompany.com"
	client := chatwork.NewChatworkClientWithDetailed(`your-api-key-here`, client, alt_host)
	...
}

```

## Post Message ##

```go
import (
	"fmt"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

func main() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	room_id := 1
	params := chatwork.RoomMessagePostParam{Body: "Hi there.", SelfUnread: optional.BoolTrue()}

	if r, err := client.Rooms().Message().Post(room_id, params); err == nil {
		fmt.Println(r.MessageId)
	} else {
		panic(err)
	}

}

```

## Create Task ##

```go
import (
	"fmt"
	"time"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

func main() {

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
```

## Get Room Messages ##

```go
import (
	"fmt"

	"github.com/magicdrive/go-chatwork"
	"github.com/magicdrive/go-chatwork/api/param/optional"
)

func main() {

	client := chatwork.NewChatworkClient(`your-api-key-here`)

	room_id := 1
	force_flang := optional.BoolTrue()

	mlist, err := client.Rooms().Message().List(room_id, force_flang)
	if err != nil {
		panic(err)
	}

	for _, item := range mlist {
		fmt.Printf("Account: '%v'\nBody: '%v'\n\n}", item.Account.Name, item.Body)
	}

}
```

## License

MIT License

Copyright (c) 2021 Hiroshi IKEGAMI

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
