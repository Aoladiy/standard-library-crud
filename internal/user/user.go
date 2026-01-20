package user

type User struct {
	Id          int     `json:"id"`
	Email       string  `json:"email"`
	FullName    *string `json:"fullName,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Age         *int    `json:"age" validate:"omitempty,gte=18"`
}
