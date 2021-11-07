package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type KeystoneClient struct {
	endpoint      string
	auth          *Auth
	token         string
	expireAt      int64
	tokenDuration int64
}

func NewKeystoneClient(host string, port int, auth *Auth) *KeystoneClient {
	return &KeystoneClient{
		endpoint:      "http://" + host + ":" + strconv.Itoa(port),
		auth:          auth,
		token:         "",
		expireAt:      0,
		tokenDuration: 82800, // The default is 23 hours
	}
}

func (c *KeystoneClient) SetTokenDuration(duration int64) error {
	if duration < 60 {
		return fmt.Errorf("token duration is less than 1 minute")
	}
	c.tokenDuration = duration
	return nil
}

func (c *KeystoneClient) Token() (string, error) {
	current := time.Now().Unix()
	if c.expireAt > current {
		return c.token, nil
	}

	// Request a new token
	jsonBytes, err := c.auth.JsonBytes()
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.endpoint+"/v2.0/tokens", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 203 {
		return "", fmt.Errorf(resp.Status)
	}
	defer resp.Body.Close()

	var result []byte
	resp.Body.Read(result)

	var tokenResult map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&tokenResult)
	if err != nil {
		return "", err
	}

	token, ok := tokenResult["access"].(map[string]interface{})["token"].(map[string]interface{})["id"].(string)
	if !ok {
		return "", fmt.Errorf("%v", tokenResult)
	}

	c.token = token
	c.expireAt = time.Now().Unix() + c.tokenDuration
	return c.token, nil
}
