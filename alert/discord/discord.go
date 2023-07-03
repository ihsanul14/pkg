package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ihsanul14/pkg/alert"
)

type DiscordAlert struct {
	URL     string
	Content ContentData
	Proxy   []string
}

type ContentData struct {
	Name    string
	Message error
	Roles   []string
}

func NewDiscordAlert() alert.Alert {
	return &DiscordAlert{}
}

func (s *DiscordAlert) BuildContent() string {
	res := fmt.Sprintf("%s ```%s```", s.Content.Name, s.Content.Message.Error())
	if s.Content.Roles != nil {
		var roles string
		for k, v := range s.Content.Roles {
			if k > 0 {
				roles = roles + fmt.Sprintf(" <@&%s>", v)
			} else {
				roles = fmt.Sprintf("<@&%s>", v)
			}
		}
		res = fmt.Sprintf("%s %s", res, roles)
	}
	return res
}

func (s *DiscordAlert) Send() error {
	if s.Content.Message != nil {
		content := map[string]string{
			"content": s.BuildContent(),
		}
		payload, err := json.Marshal(content)
		if err != nil {
			return err
		}
		req, err := http.NewRequest(alert.HTTP_POST, s.URL, bytes.NewBuffer(payload))
		if err != nil {
			return err
		}
		req.Header.Set(alert.ContentType, alert.ApplicationJson)

		client := alert.GenerateClient(s.Proxy)
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
	}
	return nil
}
