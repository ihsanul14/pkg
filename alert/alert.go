package alert

type Alert struct {
	URL     string
	Content string
	Proxy   []string
}

type Worker interface {
	Send()
}
