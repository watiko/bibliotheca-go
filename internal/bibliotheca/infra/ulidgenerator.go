package infra

import (
	"io"
	"math/rand"
	"sync"

	"github.com/oklog/ulid/v2"
)

var _ ULIDGenerator = &ulidGenerator{}

type ULIDGenerator interface {
	// len(string) must be 26
	Generate() (string, error)
	MustGenerate() string
}

type ulidGenerator struct {
	sync.Mutex
	t TimeProvider
	r io.Reader
}

func NewULIDGenerator(t TimeProvider) ULIDGenerator {
	return &ulidGenerator{
		t: t,
		r: rand.New(rand.NewSource(t.Now().UTC().UnixNano())),
	}
}

func (g *ulidGenerator) Generate() (string, error) {
	g.Lock()
	id, err := ulid.New(ulid.Timestamp(g.t.Now()), g.r)
	g.Unlock()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (g *ulidGenerator) MustGenerate() string {
	id, err := g.Generate()
	if err != nil {
		panic(err)
	}
	return id
}
