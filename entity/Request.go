package entity

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Request struct {
	ProviderAddress common.Address `json:"provider_address"`
	Value           *big.Int       `json:"value"`
	Description     string         `json:"description"`
}
