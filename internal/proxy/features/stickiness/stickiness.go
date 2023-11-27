package stickiness

import "time"

type Type string

const (
	CookieStickinessType   Type = "cookie"
	SourceIpStickinessType      = "source_ip"
)

type Stickiness struct {
	// Duration is the expiration time of the cookie
	Duration time.Duration

	// Enabled activates the feature
	Enabled bool

	// Type defines which stickiness method to use
	Type Type
}
