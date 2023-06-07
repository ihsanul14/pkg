package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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
		req, err := http.NewRequest("POST", s.URL, bytes.NewBuffer(payload))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		client := s.GenerateClient()
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return nil
	}
	return nil
}

func (s *DiscordAlert) GenerateClient() *http.Client {
	if s.Proxy == nil {
		return &http.Client{}
	}

	proxyURLs := s.Proxy
	proxyFunc := func(req *http.Request) (*url.URL, error) {
		if len(proxyURLs) == 0 {
			return nil, fmt.Errorf("no proxies available")
		}

		proxyURL := proxyURLs[0]
		proxyURLs = proxyURLs[1:]

		url, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}
		return url, nil
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxyFunc,
		},
	}

	return client
}
