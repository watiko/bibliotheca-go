package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // driver
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestPostgres struct {
	defaultDB           *sql.DB
	newConnectionString func(database string) string
	Teardown            func()
}

type DBx struct {
	*sqlx.DB
	Teardown func()
}

func NewTestPostgres() (*TestPostgres, error) {
	user := "test"
	pass := "test"
	database := "test"
	port, err := nat.NewPort("tcp", "5432")
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "postgres:11.7-alpine",
		Env: map[string]string{
			"POSTGRES_USER":     user,
			"POSTGRES_PASSWORD": pass,
			"POSTGRES_DB":       database,
		},
		ExposedPorts: []string{port.Port()},
		WaitingFor: wait.ForAll(
			wait.ForListeningPort(port),
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(60*time.Second),
		),
		AutoRemove: true,
	}
	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("could not start postgres postgresContainer: %w", err)
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get postger postgresContainer's host: %w", err)
	}

	mappedPort, err := postgresContainer.MappedPort(ctx, port)
	if err != nil {
		return nil, fmt.Errorf("could not get postger postgresContainer's port: %w", err)
	}

	newConnectionString := func(database string) string {
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, mappedPort.Int(), database)
	}

	db, err := connectDB(newConnectionString(database))
	if err != nil {
		return nil, err
	}

	teardown := func() {
		_ = db.Close()
		if err := postgresContainer.Terminate(ctx); err != nil {
			log.Fatalf("Could not stop postgres container: %v", err)
		}
	}

	postgres := &TestPostgres{
		defaultDB:           db,
		newConnectionString: newConnectionString,
		Teardown:            teardown,
	}

	return postgres, nil
}

func connectDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to the test postgres container with %s: %w", connectionString, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping failed with %s: %w", connectionString, err)
	}

	return db, nil
}

func (p TestPostgres) NewDBx() *DBx {
	db, err := p.newDB()
	if err != nil {
		log.Fatal(err)
	}
	return &DBx{
		DB: sqlx.NewDb(db.db, "postgres"),
		Teardown: func() {
			_ = db.teardown
		},
	}
}

type mydb struct {
	db       *sql.DB
	teardown func()
}

func (p TestPostgres) newDB() (*mydb, error) {
	newDatabaseName := RandString(10)

	if _, err := p.defaultDB.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", newDatabaseName)); err != nil {
		return nil, fmt.Errorf("failed to craete database: %w", err)
	}
	teardown := func() {
		_, _ = p.defaultDB.Exec(fmt.Sprintf("DROP DATABASE \"%s\"", newDatabaseName))
	}

	db, err := connectDB(p.newConnectionString(newDatabaseName))
	if err != nil {
		return nil, err
	}

	schema, err := ioutil.ReadFile("../../../../scripts/schema.sql")
	if err != nil {
		return nil, fmt.Errorf("cannot read the schema file: %w", err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return &mydb{
		db:       db,
		teardown: teardown,
	}, nil
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
