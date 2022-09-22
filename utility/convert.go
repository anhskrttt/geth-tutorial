package utility

import (
	"context"
	"geth-tutorial/initials"
	"log"
	"math"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

func WeiToEth(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	bnbValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return bnbValue
}

func IsValidAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	return re.MatchString(address)
}

func IsContractAddress(address string) bool {
	// 0x Protocol Token (ZRX) smart contract address
	bytecode, err := initials.Client.CodeAt(context.Background(), common.HexToAddress(address), nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	return isContract
}
