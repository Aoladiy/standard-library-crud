package user

import (
	"fmt"
	"sync"
)

var Users UsersStore = UsersStore{}

type User struct {
	Email       string  `json:"email"`
	FullName    *string `json:"fullName,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Age         *int    `json:"age" validate:"omitempty,gte=18"`
}
type UsersStore struct {
	users []User
	m     sync.RWMutex
}

func (u *UsersStore) GetUsers() []User {
	u.m.RLock()
	defer u.m.RUnlock()
	users := make([]User, len(u.users))
	copy(users, u.users)
	return users
}

func (u *UsersStore) GetUserById(id int) (User, error) {
	u.m.RLock()
	defer u.m.RUnlock()
	if id >= len(u.users) || id < 0 {
		return User{}, fmt.Errorf("no user with id %v", id)
	}
	user := u.users[id]
	return user, nil
}

func (u *UsersStore) AddUser(user User) {
	u.m.Lock()
	defer u.m.Unlock()
	u.users = append(u.users, user)
}

func (u *UsersStore) UpdateUser(id int, user User) (User, error) {
	u.m.Lock()
	defer u.m.Unlock()
	if id >= len(u.users) {
		return User{}, fmt.Errorf("no user with id %v", id)
	}
	oldUser := &u.users[id]
	oldUser.Email = user.Email
	if user.FullName != nil {
		oldUser.FullName = user.FullName
	}
	if user.PhoneNumber != nil {
		oldUser.PhoneNumber = user.PhoneNumber
	}
	return *oldUser, nil
}

func (u *UsersStore) DeleteUser(id int) error {
	u.m.Lock()
	defer u.m.Unlock()
	if id >= len(u.users) {
		return fmt.Errorf("no user with id %v", id)
	}
	u.users = append(u.users[:id], u.users[id+1:]...)
	return nil
}
