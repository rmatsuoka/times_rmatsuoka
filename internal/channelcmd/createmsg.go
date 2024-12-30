package channelcmd

import (
	"context"
	"fmt"

	"github.com/rmatsuoka/times_rmatsuoka/internal/channels"
	"github.com/rmatsuoka/times_rmatsuoka/internal/currnet"
	"github.com/rmatsuoka/times_rmatsuoka/internal/repository"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

type CreatingMessage struct {
	Text string
}

func (c *CreatingMessage) MessageText() string { return c.Text }

func (c *Command) CreateMessage(ctx context.Context, db xsql.DB, channelCode string, message *CreatingMessage) (channels.MessageID, error) {
	uid := currnet.UserID(ctx)

	vmessage, err := channels.ValidateCreatingMessage(message)
	if err != nil {
		return nil, err
	}

	var mid channels.MessageID
	err = db.WithinTx(ctx, func(ctx context.Context, tx xsql.Tx) error {
		cid, err := c.repositories().Channels.ID(ctx, tx, channelCode)
		if err != nil {
			return err
		}

		cuser, err := c.channelMember(ctx, tx, cid, uid)
		if err != nil {
			return err
		}

		if channels.CanCreateMessage(cuser) {
			return fmt.Errorf("no permission to post message")
		}

		mid, err = c.repositories().Messages.Create(ctx, tx, &repository.CreatingMessage{
			Message: vmessage,
			User:    uid,
			Channel: cid,
		})
		return err
	})
	return mid, err
}

func (c *Command) channelMember(ctx context.Context, db xsql.Querier, cid channels.ID, uid users.ID) (*channels.Member, error) {
	memberID, err := c.repositories().Channels.Member(ctx, db, cid, uid)
	if err != nil {
		return nil, err
	}

	user, err := c.repositories().Users.Get(ctx, db, uid)
	if err != nil {
		return nil, err
	}

	return channels.MemberFromID(user, memberID), nil
}

func CreateMessage(ctx context.Context, db xsql.DB, channelCode string, message *CreatingMessage) (channels.MessageID, error) {
	return Default.CreateMessage(ctx, db, channelCode, message)
}
