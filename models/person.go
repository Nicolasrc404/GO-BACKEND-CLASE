package models

import "time"

type Person struct {
	ID        int
	Name      string
	Age       int
	CreatedAt time.Time
}
