package v2_test

import (
	"testing"

	v2 "github.com/cqroot/go-keystone/v2"
)

var (
	client      *v2.KeystoneClient
	adminClient *v2.AdminKeystoneClient
)

func init() {
	auth := v2.NewAuth("test_tenant", "test_user", "password")
	client = v2.NewKeystoneClient("keystone.test", 5000, auth)
}

func TestToken(t *testing.T) {
	token, err := client.Token()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Token: %s", token)
}
