package messaging

const (
	// HipChatAdapter represents adapter for Bitbucket Cloud
	HipChatAdapter int = iota
)

type (
	// Controller interface for interacting with version control system
	Controller interface {
		Notify(message string, room string) error
	}

	// Config config necessary to configure messaging client
	Config struct {
		clientID     string
		clientSecret string
		AdapterID    int
		url          string
	}
)
