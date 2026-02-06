package auth

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Client struct {
	BaseURL string
	HTTP    *http.Client
}

func (c *Client) originURL() *url.URL {
	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil
	}
	u.Path = "/"
	return u
}

func NewClient(baseURL string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURL: baseURL,
		HTTP: &http.Client{
			Jar: jar,
		},
	}, nil
}

func (c *Client) CSRFToken() string {
	u := c.originURL()
	if u == nil {
		return ""
	}

	for _, cookie := range c.HTTP.Jar.Cookies(u) {
		if cookie.Name == "csrfToken" {
			return cookie.Value
		}
	}

	return ""
}

func (c *Client) DebugCookies() {
	u := c.originURL()
	if u == nil {
		log.Println("invalid backend url")
		return
	}

	log.Println("Cookies for", u.String())
	for _, cookie := range c.HTTP.Jar.Cookies(u) {
		log.Printf("  %s=%s; Path=%s\n", cookie.Name, cookie.Value, cookie.Path)
	}
}
