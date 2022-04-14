package model

import (
	"time"
)

type User struct {
	ID        uint64
	Name      string
	Password  string
	Count     uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
