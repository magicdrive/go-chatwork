package api

import "github.com/magicdrive/go-chatwork/api/my"

type MyResource struct {
	ResourceName string
	Credential   string
}

func My(credential string) MyResource {
	data := MyResource{
		ResourceName: `/my`,
		Credential:   credential,
	}
	return data
}

func (c MeResource) Tasks() my.TasksResource {
	return my.NewTasks(c)
}
