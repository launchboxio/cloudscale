package sdk

import "github.com/go-resty/resty/v2"

type Client struct {
	http *resty.Client
}

func New(endpoint string) *Client {
	client := resty.New()
	client.SetBaseURL(endpoint)
	return &Client{client}
}

func (c *Client) Certificates() *Certificates {
	return &Certificates{c}
}

func (c *Client) Listeners() *Listeners {
	return &Listeners{c}
}

func (c *Client) TargetGroups() *TargetGroups {
	return &TargetGroups{c}
}
