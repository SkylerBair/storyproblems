package usercache

import (
	"errors"
	"time"
)

// our main resource
var userMap map[string]*User

type User struct {
	username string
	createAt time.Time
}

type Storable interface {
	Get(key string) (*User, error)
	Set(key string, user *User) error
}

type storable struct {
	lastCreated time.Time
}

func New() Storable {
	s := &storable{}
	s.getLastCreated()
	return s
}

func (s *storable) getLastCreated() time.Time {
	return s.lastCreated
}

func (s *storable) Get(key string) (*User, error) {
	return nil, errors.New("torable{.Get}not impl")
}

func (s *storable) Set(key string, user *User) error {
	return errors.New("not impl")
}

// don't expose the resources directly, only through our own functions
