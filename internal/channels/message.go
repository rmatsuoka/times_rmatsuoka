package channels

import (
	"time"

	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
)

type MessageID interface {
	String() string
	MessageID()
}

type Message struct {
	ID         MessageID
	User       users.User
	Channel    ID
	Text       string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Modified   bool
}

type CreatingMessage interface {
	MessageText() string
}

type ValidCreatingMessage struct {
	CreatingMessage
}

func ValidateCreatingMessage(c CreatingMessage) (ValidCreatingMessage, error) {
	return ValidCreatingMessage{c}, nil
}

func CanCreateMessage(cuser *ChannelUser) bool {
	return cuser.Role == RoleAdmin || cuser.Role == RoleWriter
}
