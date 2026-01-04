package item

type User struct {
	Email       string  `json:"email"`
	FullName    *string `json:"fullName,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

var users []User
