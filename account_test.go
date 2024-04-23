package assets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAccount(t *testing.T) {
	tests := []struct {
		address     string
		name        string
		requireMemo bool
		workChain   int32
	}{
		{"0:5f00decb7da51881764dc3959cec60609045f6ca1b89e646bde49d492705d77f", "OKX", true, 0},
		{"EQBfAN7LfaUYgXZNw5Wc7GBgkEX2yhuJ5ka95J1JJwXXf4a8", "OKX", true, 0},
		{"0:a3935861f79daf59a13d6d182e1640210c02f98e3df18fda74b8f5ab141abf18", "Getgems Marketplace", false, 0},
		{"EQCjk1hh952vWaE9bRguFkAhDAL5jj3xj9p0uPWrFBq_GEMS", "Getgems Marketplace", false, 0},
		{"EQCjk1hh952vWaE9bRguFkAhDAL5jj3xj9p0uPWrFBq_GEMS", "Getgems Marketplace", false, 0},
		{"-1:1189458eea400d0c5dc5b1a22eda8dd009baba5465b2a99c5145733c07d9916c", "HB Pool #1", false, -1},
	}

	for _, test := range tests {
		account := GetAccountByAddress(test.address)
		assert.NotNil(t, account)
		assert.Equal(t, test.name, account.Name)
	}
}
