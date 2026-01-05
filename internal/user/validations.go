package user

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func validateUser(user User) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(user)
	if err != nil {
		log.Println("User failed validation", err)
		return err
	}
	return nil
}
