package discord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Send(t *testing.T) {
	alert := &DiscordAlert{
		URL:     "https://discord.com",
		Content: "Test Lib",
		Proxy:   nil,
	}
	t.Run("Test Send", func(t *testing.T) {
		err := alert.Send()
		assert.Nil(t, err)
	})
}
