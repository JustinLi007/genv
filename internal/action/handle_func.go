package action

import (
	"fmt"
)

type ActionHandler interface {
	Perform(ar *ActionRequest)
}

type ActionHandlerFunc func(ar *ActionRequest)

func (h ActionHandlerFunc) Perform(ar *ActionRequest) {
	h(ar)
}

func One(next ActionHandler) ActionHandler {
	return ActionHandlerFunc(func(ar *ActionRequest) {
		fmt.Println("one before")
		next.Perform(ar)
		fmt.Println("one after")
	})
}
func Two(next ActionHandler) ActionHandler {
	return ActionHandlerFunc(func(ar *ActionRequest) {
		fmt.Println("two before")
		next.Perform(ar)
		fmt.Println("two after")
	})
}
