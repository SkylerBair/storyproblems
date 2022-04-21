package usercache

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	CreatedAt time.Time
}

type Storable interface {
	Get(key string) (*User, error)
	Set(key string, user *User) error
}

type storable struct {
	lastCreated time.Time
	userMap     map[string]*User
}

func New() Storable {
	s := &storable{
		userMap: make(map[string]*User),
	}
	return s
}

func (s *storable) getLastCreated() time.Time {
	return s.lastCreated
}

func (s *storable) Get(key string) (*User, error) {
	v, exists := s.userMap[key]
	if !exists {
		return nil, fmt.Errorf("%v no user found", key)
	}

	return v, nil
}

func (s *storable) Set(key string, user *User) error {
	if err := validateKey(key); err != nil {
		return err
	}

	s.userMap[key] = user
	return nil
}

// disallowed tracks usersnames that are not allowed in our system
var disallowed map[string]interface{}

func validateKey(key string) error {
	if len(key) > 32 {
		return errors.New("ErrKeyTooLarge")
	}
	if key == "admin" {
		return errors.New("ErrInvalid")
	}
	if _, ok := disallowed[key]; ok {
		return errors.New("ErrInvalid")
	}
	return nil
}

// don't expose the resources directly, only through our own functions
