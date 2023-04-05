package discord

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type DiscordAlert struct {
	URL     string
	Content string
	Proxy   []string
}

func (s *DiscordAlert) Send() error {
	return s.send()
}

func (s *DiscordAlert) send() error {
	content := map[string]string{
		"content": s.Content,
	}
	payload, err := json.Marshal(content)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", s.URL, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
