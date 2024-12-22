package repository

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/channels"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

type Channels interface {
	ID(context.Context, xsql.Querier, string) (channels.ID, error)
	GetByCode(context.Context, xsql.Querier, string) (*channels.Channel, error)
	Create(context.Context, xsql.Tx, channels.ValidCreating) (channels.ID, error)

	AddMember(context.Context, xsql.Tx, *channels.MemberID) error
	Members(context.Context, xsql.Querier, channels.ID) ([]*channels.MemberID, error)
	Member(context.Context, xsql.Querier, channels.ID, users.ID) (*channels.MemberID, error)
}

type CreatingMessage struct {
	Message channels.ValidCreatingMessage
	User    users.ID
	Channel channels.ID
}

type Messages interface {
	ID(context.Context, xsql.Querier, string) (channels.MessageID, error)
	GetByCode(context.Context, xsql.Querier, string) (*channels.Message, error)
	Create(context.Context, xsql.Tx, *CreatingMessage) (channels.MessageID, error)
}
