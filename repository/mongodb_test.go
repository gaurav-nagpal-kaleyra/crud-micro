package repository

import (
	"firstExercise/config"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

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
		os.Setenv("MONGO_HOST", "127.0.0.1")
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

// func TestFetchBtlDataForDumping(t *testing.T) {
// 	mdbClient := MongoRepository{
// 		Client: config.Client,
// 	}

//		t.Run("Test FetchBtlDataForDumping - pass", func(t *testing.T) {
//			now := time.Now().UTC()
//			startDate := time.Date(now.Year(), now.Month(), now.Day(), 00, 00, 00, 0, time.UTC).AddDate(0, 0, -1)
//			endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC).AddDate(0, 0, -1)
//			testCompanyId := "test-company"
//			res, err := mdbClient.FetchBtlDataForDumping(startDate, endDate, testCompanyId)
//			require.NoError(t, err)
//			require.NotNil(t, res)
//		})
//	}
func TestGetBTLDocumentCountBasedOnDate(t *testing.T) {
	mongoRepo := MongoRepository{
		Client: config.Client,
	}

	t.Run("Test GetBTLDocumentCountBasedOnDate - pass", func(t *testing.T) {
		now := time.Now().UTC()
		startDate := time.Date(now.Year(), now.Month(), now.Day(), 00, 00, 00, 0, time.UTC).AddDate(0, 0, -1)
		endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC).AddDate(0, 0, -1)
		testCompanyId := "test-company"
		res, err := mongoRepo.GetBTLDocumentCountBasedOnDate(startDate, endDate, testCompanyId)
		require.NoError(t, err)
		require.GreaterOrEqual(t, res, 0)
	})
}

func TestFetchBtlDataForDumping(t *testing.T) {
	mongoRepo := MongoRepository{
		Client: config.Client,
	}

	t.Run("Test FetchBtlDataForDumping - pass", func(t *testing.T) {
		now := time.Now().UTC()
		startDate := time.Date(now.Year(), now.Month(), now.Day(), 00, 00, 00, 0, time.UTC).AddDate(0, 0, -1)
		endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC).AddDate(0, 0, -1)
		testCompanyId := "test-company"
		res, err := mongoRepo.FetchBtlDataForDumping(startDate, endDate, testCompanyId)
		require.NoError(t, err)
		require.NotNil(t, res)
	})
}
