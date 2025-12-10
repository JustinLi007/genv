package action

import (
	"fmt"
)

type ActionRequest struct {
	ex *Action
	m  map[string]any
}

func NewActionRequest(ex *Action) *ActionRequest {
	ao := &ActionRequest{
		ex: ex,
		m:  make(map[string]any),
	}
	ao.Set("action", "help")
	return ao
}

func (ar *ActionRequest) Set(key string, value any) {
	if ar.m == nil {
		return
	}
	ar.m[key] = value
}

func (ar *ActionRequest) Has(key string) bool {
	if ar.m == nil {
		return false
	}
	_, ok := ar.m[key]
	return ok
}

func (ar *ActionRequest) Get(key string) any {
	if ar.Has(key) {
		return ar.m[key]
	}
	return nil
}

func (ar *ActionRequest) Delete(key string) {
	delete(ar.m, key)
}

func (ar *ActionRequest) Size() int {
	if ar.m == nil {
		return 0
	}
	return len(ar.m)
}

func (ar *ActionRequest) Print() {
	if ar.m == nil {
		return
	}
	for k, v := range ar.m {
		fmt.Printf("%s - %v\n", k, v)
	}
}
