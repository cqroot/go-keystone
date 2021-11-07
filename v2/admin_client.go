package v2

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
)

type AdminKeystoneClient struct {
	KeystoneClient
}

func NewAdminKeystoneClient(host string, port int, auth *Auth) *AdminKeystoneClient {
	var client AdminKeystoneClient
	client.endpoint = "http://" + host + ":" + strconv.Itoa(port)
	client.auth = auth
	client.token = ""
	client.expireAt = 0
	client.tokenDuration = 82800

	return &client
}

func GetMD5Hash(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func (c *AdminKeystoneClient) ValidateToken(token string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.endpoint+"/v2.0/tokens/"+GetMD5Hash(token), nil)
	if err != nil {
		return err
	}

	auth_token, err := c.Token()
	if err != nil {
		return err
	}

	req.Header.Set("X-Auth-Token", auth_token)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 || resp.StatusCode == 203 || resp.StatusCode == 204 {
		return nil
	} else {
		return fmt.Errorf(resp.Status)
	}
}
