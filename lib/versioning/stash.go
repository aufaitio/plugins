package versioning

import (
	"net/url"

	"github.com/xoom/stash"
)

// Stash client for interfacing with Stash version control
type Stash struct {
	clientID     string
	clientSecret string
	url          string
	client       stash.Stash
	Controller
}

// NewStashClient creates a Stash client
func NewStashClient(config Config) (Stash, error) {
	client := Stash{
		clientID:     config.clientID,
		clientSecret: config.clientSecret,
		url:          config.url,
	}

	endpoint, _ := url.Parse(client.url)
	client.client = stash.NewClient(client.clientID, client.clientSecret, endpoint)

	return client, nil
}

// GetRepository clones Stash Git repository
func (stash Stash) GetRepository(project string, repository string) (stash.Repository, error) {
	return stash.client.GetRepository(project, repository)
}

// CreatePullRequest creates a pull request on remote Stash repository
func (stash Stash) CreatePullRequest(project string, repository string, title string, description string, fromRef, toRef string, reviewers []string) (stash.PullRequest, error) {
	return stash.client.CreatePullRequest(project, repository, title, description, fromRef, toRef, reviewers)
}

// GetRawFile get raw Git file
func (stash Stash) GetRawFile(project string, repository string, branch string, file string) ([]byte, error) {
	return stash.client.GetRawFile(project, repository, file, branch)
}
