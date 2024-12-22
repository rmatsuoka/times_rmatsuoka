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
	UserID     users.ID
	Channel    ID
	Text       string
	CreatedAt  time.Time
	ModifiedAt time.Time
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

func CanCreateMessage(member *Member) bool {
	return member.Role == RoleAdmin || member.Role == RoleWriter
}
