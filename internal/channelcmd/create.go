package channelcmd

import (
	"context"

	"github.com/rmatsuoka/times_rmatsuoka/internal/channels"
	"github.com/rmatsuoka/times_rmatsuoka/internal/clients"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

type Creating struct {
	Code string `json:"code"`
}

func (c *Creating) ChannelCode() string { return c.Code }

func (c *Command) Create(ctx context.Context, db xsql.DB, channel *Creating) (channels.ID, error) {
	u := clients.FromContext(ctx)

	vchannel, err := channels.ValidateCreating(channel)
	if err != nil {
		return nil, err
	}

	var cid channels.ID
	err = db.WithinTx(ctx, func(ctx context.Context, tx xsql.Tx) error {
		cid, err = c.Repository.Channels.Create(ctx, tx, vchannel)
		if err != nil {
			return err
		}
		return c.Repository.Channels.AddMember(ctx, tx, &channels.MemberID{
			Channel: cid,
			User:    u,
			Role:    channels.RoleOwenr,
		})
	})
	return cid, err
}

func Create(ctx context.Context, db xsql.DB, channel *Creating) (channels.ID, error) {
	return Default.Create(ctx, db, channel)
}
