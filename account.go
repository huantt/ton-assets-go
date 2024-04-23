package assets

import (
	_ "embed"
	"encoding/json"
	"github.com/xssnick/tonutils-go/address"
	"strings"
)

type Account struct {
	Address     string `json:"address"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	RequireMemo bool   `json:"require_memo"`
	WorkChain   int32  `json:"workchain"`
}

//go:embed data/accounts.json
var accountData []byte
var accountByAddress = make(map[string]*Account)

func init() {
	var accounts []Account
	if err := json.Unmarshal(accountData, &accounts); err != nil {
		panic(err)
	}
	for i := range accounts {
		// Convert raw address to friendly address
		addr, err := address.ParseRawAddr(accounts[i].Address)
		if err == nil {
			accounts[i].Address = addr.String()
		}
		accountByAddress[strings.ToLower(accounts[i].Address)] = &accounts[i]
	}
}

// GetAccountByAddress You can input both a friendly address and a raw address
func GetAccountByAddress(accountAddress string) *Account {
	// Try to convert raw address to friendly address
	var workChain int32
	addr, err := address.ParseRawAddr(accountAddress)
	if err == nil && addr != nil {
		accountAddress = addr.String()
		workChain = addr.Workchain()
	}
	account := accountByAddress[strings.ToLower(accountAddress)]
	if account == nil {
		return nil
	}
	account.WorkChain = workChain
	return account
}
