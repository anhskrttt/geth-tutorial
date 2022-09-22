package account

import (
	"context"
	"geth-tutorial/initials"
	"geth-tutorial/utility"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

func GetBalance(address string) *big.Float {
	// Accounts
	account := common.HexToAddress(os.Getenv("LOCAL_ADDRESS_TEST"))

	// Get balance
	balance, err := initials.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	bnbBalance := utility.WeiToEth(balance)
	return bnbBalance
}
