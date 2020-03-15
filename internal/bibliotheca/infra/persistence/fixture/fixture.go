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
	C: time.Unix(1580565602, 0).UTC(),
}

var Bookshelves = struct {
	A1      model.Bookshelf
	A2      model.Bookshelf
	Common1 model.Bookshelf
	B1      model.Bookshelf
}{
	A1:      model.Bookshelf{BookshelfID: "01E00GT5R0W3B7FT15NY84X6HD", Name: "bookshelf-a-1", CreatedAt: Times.A, UpdatedAt: Times.A},
	A2:      model.Bookshelf{BookshelfID: "01E00GT5R07R4CV6MKD83PCKPJ", Name: "bookshelf-a-2", CreatedAt: Times.A, UpdatedAt: Times.A},
	Common1: model.Bookshelf{BookshelfID: "01E00GT5R0KE0Q5WV46170EJFD", Name: "bookshelf-common-1", CreatedAt: Times.A, UpdatedAt: Times.A},
	B1:      model.Bookshelf{BookshelfID: "01E00GT6Q8XFCXSR5Y8W8ZFQTJ", Name: "bookshelf-b-1", CreatedAt: Times.B, UpdatedAt: Times.B},
}

func PrepareSeed(t *testing.T, db *sqlx.DB) {
	t.Helper()
	db.MustExec(`
DO $$
DECLARE
    time_a         timestamptz := '2020-02-01T14:00:00Z';
    time_b         timestamptz := '2020-02-01T14:00:01Z';
    time_c         timestamptz := '2020-02-01T14:00:02Z';
    alice_id       ulid        := '01E00GT5R0T01GEEH16YDGCCP4';
    bob_id         ulid        := '01E00GT6Q870BX5EQVPS8KD4AT';
    charles_id     ulid        := '01E00GT7PG35FGSG0TWZ6CDWCJ';
    g_a_id         ulid        := '01E00GT5R0NP45ST6JVZC2CB79';
    g_common_id    ulid        := '01E00GT5R01D9452B0MF62EZXA';
    g_b_id         ulid        := '01E00GT6Q82N69F6GDYQ6NDB41';
    bs_a_1_id      ulid        := '01E00GT5R0W3B7FT15NY84X6HD';
    bs_a_2_id      ulid        := '01E00GT5R07R4CV6MKD83PCKPJ';
    bs_common_1_id ulid        := '01E00GT5R0KE0Q5WV46170EJFD';
    bs_b_1_id      ulid        := '01E00GT6Q8XFCXSR5Y8W8ZFQTJ';
BEGIN
    INSERT INTO users (user_id, name, updated_at, created_at, email) VALUES
        (alice_id,   'Alice',   time_a, time_a, 'alice@a.example.com')
      , (bob_id,     'Bob',     time_b, time_b, 'bob@b.example.com')
      , (charles_id, 'Charles', time_c, time_c, 'charles@b.example.com')
      ;

    INSERT INTO groups (group_id, name, updated_at, created_at) VALUES 
        (g_a_id,      'group-a',      time_a, time_a)
      , (g_common_id, 'group-common', time_a, time_a)
      , (g_b_id,      'group-b',      time_b, time_b)
      ;

    INSERT INTO user_group_memberships (user_id, group_id) VALUES
        (alice_id,   g_a_id)
      , (alice_id,   g_common_id)
      , (bob_id,     g_common_id)
      , (bob_id,     g_b_id)
      , (charles_id, g_b_id)
      ;

    INSERT INTO bookshelves (bookshelf_id, group_id, name, updated_at, created_at) VALUES
        (bs_a_1_id,      g_a_id,      'bookshelf-a-1', time_a, time_a)
      , (bs_a_2_id,      g_a_id,      'bookshelf-a-2', time_a, time_a)
      , (bs_common_1_id, g_common_id, 'bookshelf-common-1', time_a, time_a)
      , (bs_b_1_id,      g_b_id,      'bookshelf-b-1', time_b, time_b)
      ;
END $$;
`)
}
