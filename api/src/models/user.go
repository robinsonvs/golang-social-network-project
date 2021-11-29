package models

import "time"

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Mail     string    `json:"mail,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateIn time.Time `json:"createin,omitempty"`
}
