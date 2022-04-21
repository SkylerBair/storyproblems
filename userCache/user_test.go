package usercache

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	n := New()
	want := &User{CreatedAt: time.Now()}
	err := n.Set("NewUser", want)
	if err != nil {
		t.Errorf("Failed to set user %v", err)
	}

	got, err := n.Get("NewUser")
	if want != got {
		t.Errorf("got %v; want%v *User", got, want)

		if err != nil {
			t.Errorf("Unable to get NewUser %v", got)
		}
	}

}
