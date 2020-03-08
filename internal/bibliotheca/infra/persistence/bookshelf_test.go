package persistence_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/infra/persistence"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/infra/persistence/fixture"
	"github.com/watiko/bibliotheca-go/internal/bibliotheca/types"
)

var postgres *persistence.TestPostgres

func init() {
	var err error
	postgres, err = persistence.NewTestPostgres()
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	code := m.Run()

	postgres.Teardown()

	os.Exit(code)
}

func Test_bookshelfRepo_GetAllBookshelves(t *testing.T) {
	t.Parallel()
	appCtx := types.NewTestAppContext()

	tests := []struct {
		name    string
		userID  string
		want    []*model.Bookshelf
		wantErr bool
	}{
		{name: "alice", userID: "1", want: []*model.Bookshelf{&fixture.Bookshelves.A1, &fixture.Bookshelves.A2, &fixture.Bookshelves.Common1}, wantErr: false},
		{name: "bob", userID: "2", want: []*model.Bookshelf{&fixture.Bookshelves.Common1, &fixture.Bookshelves.B1}, wantErr: false},
		{name: "charles", userID: "3", want: []*model.Bookshelf{&fixture.Bookshelves.B1}, wantErr: false},
		{name: "not_existing_user", userID: "100", want: []*model.Bookshelf{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbx := postgres.NewDBx()
			defer dbx.Teardown()
			fixture.PrepareSeed(t, dbx.DB)

			r := persistence.NewBookshelfRepository(appCtx, dbx.DB)

			got, err := r.GetAllBookshelvesForUser(context.Background(), tt.userID)
			if (err != nil) != tt.wantErr {
				assert.NoError(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
