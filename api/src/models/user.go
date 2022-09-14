package models

import (
	"errors"
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

	if step == "register" && user.Password == "" {
		return errors.New("Password is mandatory")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Mail = strings.TrimSpace(user.Mail)
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
	return nil
}
