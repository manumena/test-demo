package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Login(username, password string) error {
	payload := LoginRequest{
		Username: username,
		Password: password,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/auth/login", c.BaseURL),
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed: status %d", resp.StatusCode)
	}

	// Body intentionally ignored: tokens are in cookies
	return nil
}