package fixture

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/watiko/bibliotheca-go/internal/bibliotheca/domain/model"
)

var Times = struct {
	A time.Time
	B time.Time
	C time.Time
}{
	// timestamptz
	A: time.Unix(1580565600, 0).UTC(),
	B: time.Unix(1580565601, 0).UTC(),
	C: time.Unix(1580565601, 0).UTC(),
}

var Bookshelves = struct {
	A1      model.Bookshelf
	A2      model.Bookshelf
	Common1 model.Bookshelf
	B1      model.Bookshelf
}{
	A1:      model.Bookshelf{BookshelfID: "1", Name: "bookshelf-a-1", CreatedAt: Times.A, UpdatedAt: Times.A},
	A2:      model.Bookshelf{BookshelfID: "2", Name: "bookshelf-a-2", CreatedAt: Times.A, UpdatedAt: Times.A},
	Common1: model.Bookshelf{BookshelfID: "3", Name: "bookshelf-common-1", CreatedAt: Times.A, UpdatedAt: Times.A},
	B1:      model.Bookshelf{BookshelfID: "4", Name: "bookshelf-b-1", CreatedAt: Times.B, UpdatedAt: Times.B},
}

func PrepareSeed(t *testing.T, db *sqlx.DB) {
	t.Helper()
	db.MustExec(`
DO $$
DECLARE
    time_a timestamptz := '2020-02-01T14:00:00Z';
    time_b timestamptz := '2020-02-01T14:00:01Z';
    time_c timestamptz := '2020-02-01T14:00:02Z';
BEGIN
    INSERT INTO users (user_id, name, updated_at, created_at, email) VALUES
        (1, 'Alice',   time_a, time_a, 'alice@a.example.com')
      , (2, 'Bob',     time_b, time_b, 'bob@b.example.com')
      , (3, 'Charles', time_c, time_c, 'charles@b.example.com')
      ;
    PERFORM setval('users_user_id_seq', (SELECT max(user_id) FROM users));

    INSERT INTO groups (group_id, name, updated_at, created_at) VALUES 
        (1, 'group-a',      time_a, time_a)
      , (2, 'group-common', time_a, time_a)
      , (3, 'group-b',      time_b, time_b)
      ;
    PERFORM setval('groups_group_id_seq', (SELECT max(group_id) FROM groups));

    INSERT INTO user_group_memberships (user_id, group_id) VALUES
        (1, 1) -- alice, a
      , (1, 2) -- alice, common
      , (2, 2) -- bob, common
      , (2, 3) -- bob, b
      , (3, 3) -- charles, c
      ;

    INSERT INTO bookshelves (bookshelf_id, group_id, name, updated_at, created_at) VALUES
        (1, 1, 'bookshelf-a-1', time_a, time_a)
      , (2, 1, 'bookshelf-a-2', time_a, time_a)
      , (3, 2, 'bookshelf-common-1', time_a, time_a)
      , (4, 3, 'bookshelf-b-1', time_b, time_b)
      ;
    PERFORM setval('book_shelves_book_shelf_id_seq', (SELECT max(bookshelf_id) FROM bookshelves));
END $$;
`)
}
