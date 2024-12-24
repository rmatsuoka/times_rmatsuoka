package userinfra

import (
	"context"
	"testing"

	"github.com/rmatsuoka/times_rmatsuoka/internal/infra/testdb"
	"github.com/rmatsuoka/times_rmatsuoka/internal/repository/repositorytest"
)

func TestUsers(t *testing.T) {
	err := repositorytest.TestUsers(context.Background(), Users{}, testdb.Open())
	if err != nil {
		t.Error(err)
	}
}
