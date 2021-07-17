package goobserver

import "sync"

type Observer interface {
	Handle(arg interface{})
}

type observable struct {
	observers []Observer
	mu        *sync.Mutex
}

func NewObservable() *observable {
	return &observable{
		mu: &sync.Mutex{},
	}
}

func (l *observable) Attach(o Observer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.observers = append(l.observers, o)
}

func (l *observable) Detach(o Observer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i, v := range l.observers {
		if v == o {
			l.observers = append(l.observers[:i], l.observers[i+1:]...)
			return
		}
	}
}

func (l *observable) Notify(args interface{}) {
	for i := 0; i < len(l.observers); i++ {
		l.observers[i].Handle(args)
	}
}
