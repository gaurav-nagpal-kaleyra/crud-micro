package repository

import (
	"firstExercise/config"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
)

func TestAddUserInDB(t *testing.T) {

}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	environVars := []string{
		"MONGO_INITDB_ROOT_USERNAME=root",
		"MONGO_INITDB_ROOT_PASSWORD=root",
	}

	resource, err := pool.Run("mongo", "5.0", environVars)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		os.Setenv("MONGO_HOST","127.0.0.1")
		os.Setenv("MONGO_PORT", resource.GetPort("27017/tcp"))

		return config.MongoDBConnection()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}
