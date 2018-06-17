package messaging

import (
	"net/url"

	"github.com/tbruyelle/hipchat-go/hipchat"
)

// HipChat plugin for interfacing with HipChat client
type HipChat struct {
	client       *hipchat.Client
	clientID     string
	clientSecret string
	url          string
	Controller
}

// NewHipChatClient creates a hipchat client
func NewHipChatClient(config Config) (HipChat, error) {
	client := HipChat{
		clientID:     config.clientID,
		clientSecret: config.clientSecret,
		url:          config.url,
	}

	client.client = hipchat.NewClient(client.clientID)
	client.client.BaseURL, _ = url.Parse(client.url)

	return client, nil
}

// Notify sends message to particular HipChat room
func (chatClient HipChat) Notify(message string, room string) error {
	// TODO message formatting.
	req := hipchat.NotificationRequest{
		Color:         "purple",
		Message:       message,
		MessageFormat: "html",
		Notify:        true,
	}
	_, err := chatClient.client.Room.Notification(room, &req)

	return err
}
