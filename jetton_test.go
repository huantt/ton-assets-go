package assets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetJettonBySymbol(t *testing.T) {
	tests := []struct {
		symbol  string
		address string
	}{
		{"fish", "EQATcUc69sGSCCMSadsVUKdGwM1BMKS-HKCWGPk60xZGgwsK"},
		{"FISH", "EQATcUc69sGSCCMSadsVUKdGwM1BMKS-HKCWGPk60xZGgwsK"},
		{"Fish", "EQATcUc69sGSCCMSadsVUKdGwM1BMKS-HKCWGPk60xZGgwsK"},
	}

	for _, test := range tests {
		jetton := GetJettonBySymbol(test.symbol)
		assert.NotNil(t, jetton)
		assert.Equal(t, test.address, jetton.Address)
	}
}

func TestGetJettonByAddress(t *testing.T) {
	tests := []struct {
		address string
		symbol  string
	}{
		{"0:1371473af6c19208231269db1550a746c0cd4130a4be1ca09618f93ad3164683", "FISH"},
		{"EQATcUc69sGSCCMSadsVUKdGwM1BMKS-HKCWGPk60xZGgwsK", "FISH"},
	}

	for _, test := range tests {
		jetton := GetJettonByAddress(test.address)
		assert.NotNil(t, jetton)
		assert.Equal(t, test.symbol, jetton.Symbol)
	}
}
