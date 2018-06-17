package plugins

import (
	"fmt"

	"github.com/aufaitio/plugins/lib/messaging"
	"github.com/aufaitio/plugins/lib/versioning"
)

// NewMessaging creates messaging interface
func NewMessaging(messagingConfig messaging.Config) (messaging.Controller, error) {
	switch messagingConfig.ServiceName {
	case "hipchat":
		return messaging.NewHipChatClient(messagingConfig)
	}

	return nil, fmt.Errorf("Unsupported messaging plugin: %s", messagingConfig.ServiceName)
}

// NewVersionControl creates notifier
func NewVersionControl(versionConfig versioning.Config) (versioning.Controller, error) {
	switch versionConfig.ServiceName {
	case "stash":
		return versioning.NewStashClient(versionConfig)
	}

	return nil, fmt.Errorf("Unsupported version control plugin: %s", versionConfig.ServiceName)
}
