package models

import (
	"api/src/safety"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Mail     string    `json:"mail,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateIn time.Time `json:"createin,omitempty"`
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name is mandatory")
	}

	if user.Nick == "" {
		return errors.New("Nick is mandatory")
	}

	if user.Mail == "" {
		return errors.New("Mail is mandatory")
	}

	if err := checkmail.ValidateFormat(user.Mail); err != nil {
		return errors.New("The email is inválid")
	}

	if step == "register" && user.Password == "" {
		return errors.New("Password is mandatory")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Mail = strings.TrimSpace(user.Mail)

	if step == "register" {
		passwordWithHash, err := safety.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordWithHash)
	}

	return nil
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}
	return nil
}
