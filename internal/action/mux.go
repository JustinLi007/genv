package action

import (
	"github.com/JustinLi007/genv/internal/assert"
)

type MethodTable struct {
	table map[string]*HandlerTable
}

func NewMethodTable() *MethodTable {
	mt := &MethodTable{
		table: make(map[string]*HandlerTable),
	}
	return mt
}

func (mt *MethodTable) Set(key string) {
	if mt.table == nil {
		return
	}
	if _, ok := mt.table[key]; ok {
		return
	}
	mt.table[key] = NewHandlerTable()
}

func (mt *MethodTable) Has(key string) bool {
	if mt.table == nil {
		return false
	}
	_, ok := mt.table[key]
	return ok
}

func (mt *MethodTable) Get(key string) *HandlerTable {
	if mt.Has(key) {
		return mt.table[key]
	}
	return nil
}

func (mt *MethodTable) Delete(key string) {
	delete(mt.table, key)
}

type HandlerTable struct {
	table map[string]ActionHandlerFunc
}

func NewHandlerTable() *HandlerTable {
	ht := &HandlerTable{
		table: make(map[string]ActionHandlerFunc),
	}
	return ht
}

func (ht *HandlerTable) Set(key string, value ActionHandlerFunc) {
	if ht.table == nil {
		return
	}
	ht.table[key] = value
}

func (ht *HandlerTable) Has(key string) bool {
	if ht.table == nil {
		return false
	}
	_, ok := ht.table[key]
	return ok
}

func (ht *HandlerTable) Get(key string) ActionHandlerFunc {
	if ht.Has(key) {
		return ht.table[key]
	}
	return nil
}

func (ht *HandlerTable) Delete(key string) {
	delete(ht.table, key)
}

type Mux struct {
	routes *MethodTable
}

func NewMux() *Mux {
	mux := &Mux{
		routes: NewMethodTable(),
	}
	return mux
}

func (mx *Mux) Register(method, action string, handler ActionHandlerFunc) {
	if mx.routes == nil {
		return
	}

	mx.routes.Set(method)
	handlerTable := mx.routes.Get(method)
	assert.NotNil(handlerTable, "mux: nil handler table")
	handlerTable.Set(action, handler)
}

func (mx *Mux) Perform(ar *ActionRequest) {
	method, ok := ar.Get("method").(string)
	assert.True(ok, "mux: no method specified")

	action, ok := ar.Get("action").(string)
	assert.True(ok, "mux: no action specified")

	fn := mx.routes.Get(method).Get(action)
	assert.NotNil(fn, "mux: no handler found")
	fn(ar)
}
