package github

import (
	"time"
)

// Webhook is the payload GitHub sends on a webhook event
type Webhook struct {
	Action     string     `json:"action"`
	Release    Release    `json:"release"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// Release is the data sent when the event of the payload is a repository release
type Release struct {
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	HTMLURL     string    `json:"html_url"`
	AssetsURL   string    `json:"assets_url"`
	UploadURL   string    `json:"upload_url"`
	TarballURL  string    `json:"tarball_url"`
	ZipballURL  string    `json:"zipball_url"`
	NodeID      string    `json:"node_id"`
	TagName     string    `json:"tag_name"`
	Branch      string    `json:"target_commitish"`
	Name        string    `json:"name"`
	Body        string    `json:"body"`
	Draft       bool      `json:"draft"`
	Prerelease  bool      `json:"prerelease"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
}

// Repository is the data related to a GitHub repository
type Repository struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Owner       Owner  `json:"owner"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
}

// Sender is the data related to the GitHub event creator, can be a user or a bot
type Sender struct {
	Username  string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

// Owner is the data related to the GitHub repository owner, can be a user or and organisation
type Owner struct {
	AvatarURL string `json:"avatar_url"`
}
