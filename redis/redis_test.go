package redis

import (
	"firstExercise/config"
	userModel "firstExercise/model/user"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedis(t *testing.T) {
	os.Setenv("REDIS_HOST", "localhost")
	os.Setenv("REDIS_PORT", "6379")
	os.Setenv("REDIS_DATABASE", "0")
	os.Setenv("REDIS_PASS", "")
	err := config.RedisConnection()

	require.NoError(t, err)

	mockUser := userModel.User{
		UserId:       10,
		UserName:     "testUser",
		UserAge:      20,
		UserLocation: "USA",
	}

	t.Run("Add Into Redis Cache", func(t *testing.T) {
		err := AddIntoDBRedis(&mockUser)
		require.NoError(t, err)
	})

	t.Run("Read from Redis Cache", func(t *testing.T) {
		foundUser, err := ReadFromDBRedis(strconv.Itoa(mockUser.UserId))
		require.NoError(t, err)
		require.Equal(t, mockUser, foundUser)
	})
}
