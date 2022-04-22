package jobq

import (
	"errors"
	"fmt"
)

// jobq is a job queue that is a first in, first out processor
// it is safe to use concurrently.

var jobQuantity int = 100000

type Job struct {
	Name string
	Fn   func() error
}

type Queue interface {
	Inqueue(j Job) error
	Outqueue() Job
}

type jobList struct {
	qchan chan Job
}

func Newq() *jobList {
	u := &jobList{
		qchan: make(chan Job, jobQuantity),
	}
	return u
}

func (u *jobList) Inqueue(j Job) error {
	if j.Fn == nil {
		return errors.New("no job func found")
	}
	u.qchan <- j
	return nil
}

func (u *jobList) Outqueue() Job {
	for {
		select {
		case j := <-u.qchan:
			fmt.Println("received message", j)
			return j
		default:
			fmt.Println("no message received")
		}
	}
}
