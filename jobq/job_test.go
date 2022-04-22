package jobq_test

import (
	"reflect"
	"testing"

	"github.com/SkylerBair/storyproblems/jobq"
)

func TestInQueue(t *testing.T) {
	n := jobq.Newq()
	want := jobq.Job{Fn: func() error { return nil }}
	err := n.Inqueue(want)
	if err != nil {
		t.Errorf("failed to add job to queue %v", err)
	}

}

func TestOutQueue(t *testing.T) {
	n := jobq.Newq()
	j := jobq.Job{Fn: func() error { return nil }}
	n.Inqueue(j)
	output := n.Outqueue()
	if reflect.DeepEqual(j, output) {
		t.Errorf("failed to return job %v", output)
	}
}
