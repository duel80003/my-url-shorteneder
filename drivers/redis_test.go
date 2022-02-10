package drivers

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConnection(t *testing.T) {
	ping := RedisClient.Ping()
	str, err := ping.Result()
	assert.Nil(t, err)
	assert.Equal(t, "pong", strings.ToLower(str))
}
