package v2

import "encoding/json"

type Auth struct {
	Auth struct {
		PasswordCredentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"passwordCredentials"`
		TenantName string `json:"tenantName"`
	} `json:"auth"`
}

func NewAuth(tenant string, username string, password string) (auth *Auth) {
	auth = new(Auth)
	auth.Auth.TenantName = tenant
	auth.Auth.PasswordCredentials.Username = username
	auth.Auth.PasswordCredentials.Password = password
	return
}

func (a *Auth) JsonBytes() ([]byte, error) {
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}
