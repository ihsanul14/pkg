package discord

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Send(t *testing.T) {
	var err error
	alert := &DiscordAlert{
		URL: "https://discord.com",
		Content: ContentData{
			Name:    "Test",
			Message: err,
			Roles:   []string{"917325776456663060", "1093006681425858560"},
		},
		Proxy: nil,
	}
	t.Run("Test Not Send", func(t *testing.T) {
		err := alert.Send()
		assert.Nil(t, err)
	})

	err = fmt.Errorf("%v", "test content")
	alert.Content.Message = err
	t.Run("Test Not Send", func(t *testing.T) {
		err := alert.Send()
		assert.Nil(t, err)
	})

	alert.Proxy = []string{"http://proxy:8000"}
	t.Run("Test Send With Proxy", func(t *testing.T) {
		err := alert.Send()
		assert.NotNil(t, err)
	})
}
