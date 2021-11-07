package v2_test

import (
	"testing"

	v2 "github.com/cqroot/go-keystone/v2"
)

func init() {
	adminAuth := v2.NewAuth("test_tenant", "test_admin_user", "password")
	adminClient = v2.NewAdminKeystoneClient("keystone.test", 35357, adminAuth)
}

func TestAdminToken(t *testing.T) {
	token, err := adminClient.Token()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Token: %s", token)
}

func TestValidateToken(t *testing.T) {
	token, err := client.Token()
	if err != nil {
		t.Error(err)
		return
	}

	err = adminClient.ValidateToken(token)
	if err != nil {
		t.Error(err)
	}
}
