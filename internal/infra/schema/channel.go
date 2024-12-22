package schema

import "time"

type Channel struct {
	ID        int64
	Code      string
	CreatedAt time.Time
}

type Member struct {
	ChannelID int64
	UserID    int64
	Role      int
	CreatedAt time.Time
}
