package user

import "github.com/Aoladiy/standard-library-crud/internal/db"

type User struct {
	Id          int     `json:"id"`
	Email       string  `json:"email"`
	FullName    *string `json:"fullName,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Age         *int    `json:"age" validate:"omitempty,gte=18"`
}

func GetUsers() ([]User, error) {
	return Repo{db: db.DB}.getUsers()
}

func GetUserById(id int) (user User, err error) {
	return Repo{db: db.DB}.getUserById(id)
}

func CreateUser(user User) (id int, err error) {
	return Repo{db: db.DB}.createUser(user)
}

func UpdateUser(user User) (err error, ok bool) {
	return Repo{db: db.DB}.updateUser(user)
}

func DeleteUserById(id int) (err error, ok bool) {
	return Repo{db: db.DB}.deleteUserById(id)
}
