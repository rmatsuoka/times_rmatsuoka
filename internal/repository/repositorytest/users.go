package repositorytest

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"

	"github.com/rmatsuoka/times_rmatsuoka/internal/repository"
	"github.com/rmatsuoka/times_rmatsuoka/internal/types"
	"github.com/rmatsuoka/times_rmatsuoka/internal/users"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/diff"
	"github.com/rmatsuoka/times_rmatsuoka/internal/x/xsql"
)

var testUsers = []*users.User{
	{Code: "code_0", Name: "name zero"},
	{Code: "code_1", Name: "name one"},
	{Code: "code_2", Name: "name two"},
}

func init() {
	// validate users
	for _, u := range testUsers {
		_, err := users.ValidateCreating(u)
		if err != nil {
			panic(err)
		}
	}
}

func sameUser(u, v *users.User) bool {
	return u.Code == v.Code && u.Name == v.Name
}

func TestUsers(ctx context.Context, repo repository.Users, db xsql.DB) error {
	err := db.WithinTx(ctx, func(ctx context.Context, tx xsql.Tx) error {
		const N = 10
		userMap := make(map[users.ID]*users.User)
		for _, u := range testUsers {
			id, err := repo.Create(ctx, tx, users.ValidCreating{Creating: u})
			if err != nil {
				return fmt.Errorf("repo.Create(ctx, tx, %#v) returned non-nil error: %w", u, err)
			}
			userMap[id] = u
			_, err = repo.Create(ctx, tx, users.ValidCreating{Creating: u})
			if !errors.Is(err, types.ErrExist) {
				return fmt.Errorf("repo.Create(ctx, tx, %#v) twice time returned %q as error, want %q as err", u, err, types.ErrExist)
			}
		}

		for id, u := range userMap {
			gotID, err := repo.ID(ctx, tx, u.Code)
			if err != nil {
				return fmt.Errorf("repo.ID(ctx, tx, %s) returned non-nil error: %w", u.Code, err)
			}
			if gotID != id {
				return fmt.Errorf("repo.ID(ctx, tx, %s) returned %#v, want %#v", u.Code, gotID, id)
			}

			got, err := repo.Get(ctx, tx, id)
			if err != nil {
				return fmt.Errorf("repo.Get(ctx, tx, %s) returned non-nil error: %w", id, err)
			}
			if !sameUser(got, u) {
				return fmt.Errorf("repo.Get(ctx, tx, %s) returned %#v, want %#v", id, got, u)
			}

			gotByCode, err := repo.GetByCode(ctx, tx, u.Code)
			if err != nil {
				return fmt.Errorf("repo.GetByCode(ctx, tx, %s) returned non-nil error: %w", u.Code, err)
			}
			if !sameUser(gotByCode, u) {
				return fmt.Errorf("repo.GetByCode(ctx, tx, %s) returned %#v, want %#v", u.Code, got, u)
			}
		}

		gotMap, err := repo.GetMany(ctx, tx, maps.Keys(userMap))
		if err != nil {
			return fmt.Errorf("repo.GetMany(ctx, tx, %+v) returned non-nil error: %w", slices.Collect(maps.Keys(userMap)), err)
		}

		if !maps.EqualFunc(userMap, gotMap, func(u, v *users.User) bool { return sameUser(u, v) }) {
			return fmt.Errorf("repo.GetMany(ctx, tx, %+v) returned what not expected\n%s",
				slices.Collect(maps.Keys(userMap)),
				diff.MapsFunc(userMap, gotMap,
					func(u, v *users.User) bool { return sameUser(u, v) }))
		}

		return errNoCommit
	})
	if errors.Is(err, errNoCommit) {
		return nil
	}
	return err
}
