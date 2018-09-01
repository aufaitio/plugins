package versioning

import (
	"github.com/xoom/stash"
)

const (
	// BitbucketCloudAdapter represents adapter for Bitbucket Cloud
	BitbucketCloudAdapter int = iota
	// BitbucketServerAdapter represents adapter for on premise Bitbucket (Stash)
	BitbucketServerAdapter
	// GitHubAdapter represents GitHub adapter
	GitHubAdapter
	// GitLabAdapter represents GitLab adapter
	GitLabAdapter
)

type (
	// Controller interface for interacting with version control system
	Controller interface {
		GetRepository(project string, repository string) (stash.Repository, error)
		CreatePullRequest(project string, repository string, title string, description string, fromRef, toRef string, reviewers []string) (stash.PullRequest, error)
		GetRawFile(project string, repository string, branch string, file string) ([]byte, error)
	}

	// Config config necessary to configure version control client
	Config struct {
		clientID     string
		clientSecret string
		room         string
		AdapterID    int
		url          string
	}
)
