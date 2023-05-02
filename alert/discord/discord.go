package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DiscordAlert struct {
	URL     string
	Content ContentData
	Proxy   []string
}

type ContentData struct {
	Name    string
	Message string
	Roles   []string
}

func (s *DiscordAlert) buildContent() string {
	res := fmt.Sprintf("%s ```%s```", s.Content.Name, s.Content.Message)
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
	if s.Content.Message != "" {
		return s.send()
	}
	return nil
}

func (s *DiscordAlert) send() error {
	content := map[string]string{
		"content": s.buildContent(),
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

	client := s.generateClient()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (s *DiscordAlert) generateClient() *http.Client {
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
