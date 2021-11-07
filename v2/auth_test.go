package v2_test

import (
	"encoding/json"
	"testing"

	v2 "github.com/cqroot/go-keystone/v2"
)

func TestAuth(t *testing.T) {
	auth := v2.NewAuth("tenant", "username", "password")

	jsonBytes, err := json.Marshal(auth)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(jsonBytes))
}
