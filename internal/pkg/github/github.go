package github

import (
	"time"
)

// TODO: Add comments
type Webhook struct {
	Action     string     `json:"action"`
	Release    Release    `json:"release"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}

// TODO: Add comments
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

// TODO: Add comments
type Repository struct {
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Owner       Owner  `json:"owner"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
}

// TODO: Add comments
type Sender struct {
	Username  string `json:"login"`
	AvatarURL string `json:"avatar_url"`
}

// TODO: Add comments
type Owner struct {
	AvatarURL string `json:"avatar_url"`
}
