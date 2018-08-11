package versioning

import (
	"net/url"

	"github.com/xoom/stash"
)

type (
	// Stash client for interfacing with Stash version control
	Stash struct {
		clientID     string
		clientSecret string
		url          string
		client       stash.Stash
		Controller
	}

	// StashHook webhook from Stash service
	StashHook struct {
		Changesets struct {
			Values []struct {
				Changes struct {
					Values []struct {
						Path struct {
							Components []string `json:"components"`
						} `json:"path"`
					} `json:"values"`
				} `json:"changes"`
			} `json:"values"`
		} `json:"changesets"`
		Repository struct {
			Name    string `json:"name"`
			Project struct {
				Name string `json:"name"`
			} `json:"project"`
		} `json:"repository"`
	}
)

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
