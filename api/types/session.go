package types

import "time"

type Session struct {
	ID          int       `json:"id`
	DownloadURL string    `json:"download_url"`
	Expiration  time.Time `json:"expiration"`
}
