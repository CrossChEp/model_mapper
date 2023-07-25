package model_mapper

import (
	"fmt"
	"testing"
)

type UserAddModel struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Patronymic  string `json:"patronymic"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserGetModel struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Patronymic  string `json:"patronymic"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	State       string `json:"state"`
}

func TestMapWithNullFields(t *testing.T) {
	user := UserAddModel{
		FirstName:   "adasda",
		LastName:    "asdasd",
		Patronymic:  "asdasdsa",
		PhoneNumber: "",
		Password:    "",
	}
	userGet := UserGetModel{
		LastName: "asdasd",
		Email:    "sadasdas",
	}
	if err := mapWithNullFields(&userGet, &user); err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", userGet)
}
