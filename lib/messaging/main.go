package messaging

type (
	// Controller interface for interacting with version control system
	Controller interface {
		Notify(message string, room string) error
	}

	// Config config necessary to configure messaging client
	Config struct {
		clientID     string
		clientSecret string
		ServiceName  string
		url          string
	}
)
