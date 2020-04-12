package dhsclient

type DoHOption struct {
	BaseURL string
}
type ModDoHOption func(option *DoHOption)

func WithBaseURL(s string) ModDoHOption {
	return func(option *DoHOption) {
		option.BaseURL = s
	}
}
