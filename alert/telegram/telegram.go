package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ihsanul14/pkg/alert"
)

type TelegramAlert struct {
	URL     string
	Content ContentData
	Proxy   []string
}

type ContentData struct {
	Name    string
	Message error
	ChatId  string
}

func NewTelegramAlert() alert.Alert {
	return &TelegramAlert{}
}

func (s *TelegramAlert) BuildContent() string {
	res := fmt.Sprintf("%s \n%s", s.Content.Name, s.Content.Message.Error())
	return res
}

func (s *TelegramAlert) Send() error {
	if s.Content.Message != nil {
		content := map[string]string{
			"text":    s.BuildContent(),
			"chat_id": s.Content.ChatId,
			// "parse_mode": "MarkdownV2",
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

func (s *TelegramAlert) GenerateClient() *http.Client {
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
