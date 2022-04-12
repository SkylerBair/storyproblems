package main

import (
	"fmt"
	"time"
)

type Events interface {
	Pub(e event)
	Sub() chan event
}

var _ Events = (*eventManager)(nil)

type event struct {
	message string
}

type eventManager struct {
	events chan event
}

func NewEventManager() *eventManager {
	s := &eventManager{events: make(chan event, 10000)}
	return s
}

func (e *eventManager) Pub(evt event) {
	e.events <- evt
}

func (e *eventManager) Sub() chan event {
	return e.events
}

func main() {
	MGR := NewEventManager()
	sub := MGR.Sub()
	go func() {
		for {
			select {
			case s := <-sub:
				fmt.Printf("s.message: %v\n", s.message)
			default:
				//fmt.Println("waiting")
			}
		}
	}()

	pusher := MGR.Pub
	pusher(event{
		message: "pusher was",
	})
	MGR.Pub(event{
		message: "hello",
	})
	MGR.Pub(event{
		message: "hello",
	})
	MGR.Pub(event{
		message: "hello",
	})
	MGR.Pub(event{
		message: "hello",
	})
	MGR.Pub(event{
		message: "hello",
	})
	MGR.Pub(event{
		message: "hello",
	})
	MGR.Pub(event{
		message: "hello",
	})
	MGR.Pub(event{
		message: "hello",
	})
	time.Sleep(time.Microsecond * 100)

}
