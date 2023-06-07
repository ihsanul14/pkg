package alert

import "net/http"

type Alert interface {
	Send() error
	BuildContent() string
	GenerateClient() *http.Client
}
