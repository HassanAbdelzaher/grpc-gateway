package middlewares

import (
	"errors"
	"log"
	"sync"
	"time"
)

var defaultConnectionsPerUser = 1
var mx sync.Mutex

type UserQue map[string]chan bool

func NewUserQue() UserQue {
	return make(map[string]chan bool)
}
func (u *UserQue) Set(key string, size int) {
	var x map[string]chan bool = *u
	if size == 0 {
		x[key] = nil
		return
	}
	x[key] = make(chan bool, size)
}
func (u *UserQue) get(key string) (chan bool, bool) {
	var x map[string]chan bool = *u
	v, ok := x[key]
	if v == nil {
		return nil, false
	}
	return v, ok
}
func (u *UserQue) Wait(key string) error {
	mx.Lock()
	ch, ok := u.get(key)
	if ok {
		mx.Unlock()
		select {
		case ch <- true:
			return nil
		case <-time.After(5 * time.Minute):
			return errors.New("الخادم مشغول: يوجد اعمال اخرى تخص الفرع قيد التنفيذ برجاء المحاولة مرة اخرى بعد دقيقة")
		}
	} else {
		log.Printf("cretating channel for "+key+" with length :%v", defaultConnectionsPerUser)
		u.Set(key, defaultConnectionsPerUser)
		mx.Unlock()
	}
	return nil
}

func (u *UserQue) Done(key string) {
	mx.Lock()
	defer mx.Unlock()
	ch, ok := u.get(key)
	if ok {
		select {
		case <-ch:
		default:
		}
	}
}
