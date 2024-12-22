package schema

import (
	"database/sql"
	"time"
)

type Message struct {
	ID         int64
	UserID     int64
	ChannelID  int64
	Text       string
	CreatedAt  time.Time
	ModifiedAt sql.Null[time.Time]
}
