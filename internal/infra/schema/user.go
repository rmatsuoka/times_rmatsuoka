package schema

import "time"

type User struct {
	ID        int64
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
