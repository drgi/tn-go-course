package psql

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	database "github.com/tn-go-course/go-search/hw-16/db-app/pkg"
)

var testDB *DB

const (
	dbPass = "1234"
	dbUser = "gopher"
	dbName = "films"
)

func TestMain(m *testing.M) {
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = dockerPool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := dockerPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", dbPass),
			fmt.Sprintf("POSTGRES_USER=%s", dbUser),
			fmt.Sprintf("POSTGRES_DB=%s", dbName),
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	log.Println(resource)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s/%s", dbUser, dbPass, hostAndPort, dbName)
	log.Println("Connecting to database on url: ", databaseUrl)
	resource.Expire(120)

	dockerPool.MaxWait = 120 * time.Second
	if err = dockerPool.Retry(func() error {
		testDB, err = New(databaseUrl)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	defer testDB.p.Close()

	// Добавляем схему
	b, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatalf("Could not read schema: %s", err)
	}

	_, err = testDB.p.Exec(context.Background(), string(b))
	if err != nil {
		log.Fatalf("Could not apply schema: %s", err)
	}

	b, err = os.ReadFile("data.sql")
	if err != nil {
		log.Fatalf("Could not read data sql: %s", err)
	}

	_, err = testDB.p.Exec(context.Background(), string(b))
	if err != nil {
		log.Fatalf("Could not insert data: %s", err)
	}

	code := m.Run()

	if err := dockerPool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestDB_CreateFilms(t *testing.T) {
	ctx := context.Background()
	filmUpdate := []*database.Film{
		{
			ID:          1,
			Name:        "Побег из Шоушенка",
			RealiseYear: 1995,
			Rating:      "PG-18",
			Profit:      2893458,
			StudioID:    1,
		},
	}
	err := testDB.CreateFilms(ctx, filmUpdate)
	if err != nil {
		t.Errorf("DB.Films() error = %v, wantErr %v", err, true)
		return
	}
}

func TestDB_Films(t *testing.T) {
	ctx := context.Background()
	filter := &database.FilterFilm{}
	got, err := testDB.Films(ctx, filter)
	if err != nil {
		t.Errorf("DB.Films() error = %v, wantErr %v", err, true)
		return
	}
	if len(got) != 4 {
		t.Errorf("DB.Films() got %d, want %d", len(got), 4)
	}
}

func TestDB_DeleteFilms(t *testing.T) {
	ctx := context.Background()
	err := testDB.DeleteFilm(ctx, 1)
	if err != nil {
		t.Errorf("DB.Films() error = %v, wantErr %v", err, true)
		return
	}

}

func TestDB_UpdateFilms(t *testing.T) {
	ctx := context.Background()
	filmUpdate := &database.Film{
		ID:          1,
		Name:        "Побег из Шоушенка",
		RealiseYear: 1994,
		Rating:      "PG-10",
		Profit:      2893458,
		StudioID:    1,
	}
	err := testDB.UpdateFilm(ctx, filmUpdate)
	if err != nil {
		t.Errorf("DB.Films() error = %v, wantErr %v", err, true)
		return
	}

}
