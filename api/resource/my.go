package resource

import (
	my "github.com/magicdrive/go-chatwork/api/resource/my_sub"
)

type MyResource struct {
	ResourceName string
	Credential   string
}

func NewMyResource(credential string) MyResource {
	data := MyResource{
		ResourceName: `/my`,
		Credential:   credential,
	}
	return data
}

func (c MeResource) Status() my.StatusResource {
	return my.NewStatus(c.ResourceName, c.Credential)
}

func (c MeResource) Tasks() my.TasksResource {
	return my.NewTasks(c.ResourceName, c.Credential)
}
