package role

import "time"

type Role struct {
	ID          uint
	Name        string
	Deskription string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
