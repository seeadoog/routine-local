package rlocal

import (
	"runtime"
	"bytes"
	"strconv"
	"sync"
)

var defaultLocal = NewRoutineLocal()

func DefaultRoutineLocal() IRoutineLocal {
	return defaultLocal
}

func Get(k string) interface{} {
	return defaultLocal.Get(k)
}

func Set(k string, v interface{}) () {
	defaultLocal.Set(k, v)
}

func Remove(k string) {
	defaultLocal.Remove(k)
}

func RemoveAll() {
	defaultLocal.RemoveAll()
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

type IRoutineLocal interface {
	Set(string, interface{})
	Get(string) interface{}
	Remove(string)
	RemoveAll()
	SetFin(k string, v interface{}, fin func())
}

type RoutineLocal struct {
	vals sync.Map
}

func NewRoutineLocal() IRoutineLocal {
	return &RoutineLocal{vals: sync.Map{}}
}

func (r *RoutineLocal) Set(k string, v interface{}) {
	routineNo := GetGID()
	vm, ok := r.vals.Load(routineNo)
	if !ok {
		vm = make(map[string]interface{})
		r.vals.Store(routineNo, vm)
	}
	vm.(map[string]interface{})[k] = v
}

func (r *RoutineLocal) Get(k string) interface{} {
	routineNo := GetGID()
	vm, ok := r.vals.Load(routineNo)
	if !ok {
		return nil
	}
	return vm.(map[string]interface{})[k]
}

func (r *RoutineLocal) Remove(k string) {
	routineNo := GetGID()
	vm, ok := r.vals.Load(routineNo)
	if !ok {
		return
	}
	delete(vm.(map[string]interface{}), k)
}

func (r *RoutineLocal) RemoveAll() {
	r.vals.Delete(GetGID())
}

func (r *RoutineLocal) SetFin(k string, v interface{}, fin func()) {
	r.Set(k, v)
	defer r.RemoveAll()
	fin()
}
