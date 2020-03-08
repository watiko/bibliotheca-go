package persistence_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/infra/persistence"
)

func Test_TestPostgres(t *testing.T) {
	postgres, err := persistence.NewTestPostgres()
	if !assert.NoError(t, err) {
		return
	}
	defer postgres.Teardown()

	t.Run("NewDB", func(t *testing.T) {
		db := postgres.NewDBx()

		t.Run("QueryNotExistingTable", func(t *testing.T) {
			_, err = db.Exec("select * from book")
			assert.Error(t, err)
		})

		t.Run("QueryExistingTable", func(t *testing.T) {
			_, err = db.Exec("select * from books")
			assert.NoError(t, err)
		})
	})
}
