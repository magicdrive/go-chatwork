package resource

import (
	"github.com/magicdrive/go-chatwork/api"
	my "github.com/magicdrive/go-chatwork/api/resource/my"
)

// MyResource chatwork api my resource.
type MyResource struct {
	ResourceName string
	Client       *api.ChatworkApiClient
}

// NewMyResource new chatwork api my resource.
func NewMyResource(client *api.ChatworkApiClient) MyResource {
	data := MyResource{
		ResourceName: `/my`,
		Client:       client,
	}
	return data
}

// Status new chatwork api my/status resource.
func (c MyResource) Status() my.MyStatusResource {
	return my.NewMyStatus(c.ResourceName, c.Client)
}

// Tasks new chatwork api my/tasks resource.
func (c MyResource) Tasks() my.MyTasksResource {
	return my.NewMyTasks(c.ResourceName, c.Client)
}
