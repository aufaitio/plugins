package plugins

import (
	"errors"
	"github.com/quantumew/plugins/lib/messaging"
	"github.com/quantumew/plugins/lib/versioning"
)

// NewMessaging creates messaging interface
func NewMessaging(messagingConfig messaging.Config) (messaging.Controller, error) {
	switch messagingConfig.AdapterID {
	case messaging.HipChatAdapter:
		return messaging.NewHipChatClient(messagingConfig)
	}

	return nil, errors.New("Unsupported messaging plugin")
}

// NewVersionControl creates notifier
func NewVersionControl(versionConfig versioning.Config) (versioning.Controller, error) {
	switch versionConfig.AdapterID {
	case versioning.BitbucketServerAdapter:
		return versioning.NewStashClient(versionConfig)
	}

	return nil, errors.New("Unsupported version control plugin")
}
