package domain

import "time"

type User struct {
	ID       int64
	Name     string
	Gender   Gender
	Birthday time.Time
}
