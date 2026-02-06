package students

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	BaseURL   string
	HTTP      *http.Client
	CSRFToken func() string
}

func NewClient(baseURL string, httpClient *http.Client, csrfTokenFn func() string) *Client {
	return &Client{
		BaseURL:   baseURL,
		HTTP:      httpClient,
		CSRFToken: csrfTokenFn,
	}
}

func (c *Client) GetByID(id int) (*Student, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/students/%d", c.BaseURL, id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-CSRF-Token", c.CSRFToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized: missing or invalid cookies / csrf token")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch student: status %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("RAW backend response: %s\n", string(bodyBytes))

	var dto studentDTO
	if err := json.Unmarshal(bodyBytes, &dto); err != nil {
		return nil, err
	}

	return toStudent(dto), nil
}
