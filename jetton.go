package assets

import (
	_ "embed"
	"encoding/json"
	"github.com/xssnick/tonutils-go/address"
	"strings"
)

type Jetton struct {
	Address     string   `json:"address"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Name        string   `json:"name"`
	Social      []string `json:"social"`
	Symbol      string   `json:"symbol"`
	Websites    []string `json:"websites"`
}

//go:embed data/jettons.json
var jettonsData []byte
var jettonBySymbol = make(map[string]*Jetton)
var jettonByAddress = make(map[string]*Jetton)

func init() {
	var jettons []Jetton
	if err := json.Unmarshal(jettonsData, &jettons); err != nil {
		panic(err)
	}
	for i := range jettons {
		// Convert raw address to friendly address
		addr, err := address.ParseRawAddr(jettons[i].Address)
		if err == nil {
			jettons[i].Address = addr.String()
		}
		jettonBySymbol[strings.ToLower(jettons[i].Symbol)] = &jettons[i]
		jettonByAddress[strings.ToLower(jettons[i].Address)] = &jettons[i]
	}
}

// GetJettonBySymbol symbol is case-insensitive
func GetJettonBySymbol(symbol string) *Jetton {
	return jettonBySymbol[strings.ToLower(symbol)]
}

// GetJettonByAddress You can input both a friendly address and a raw address
func GetJettonByAddress(jettonAddress string) *Jetton {
	// Try to convert raw address to friendly address
	addr, err := address.ParseRawAddr(jettonAddress)
	if err == nil && addr != nil {
		jettonAddress = addr.String()
	}
	return jettonByAddress[strings.ToLower(jettonAddress)]
}
