package smart_contracts

import (
	"fmt"
	"geth-tutorial/initials"
	"geth-tutorial/token"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func QueryERC20() {
	tokenAddress := common.HexToAddress("0xf1E790642F735a4E77A8570D91d54880FaF1aA3E")
	instance, err := token.NewToken(tokenAddress, initials.Client)
	if err != nil {
		log.Fatal(err)
	}

	_ = instance

	address := common.HexToAddress(os.Getenv("LOCAL_ADDRESS_TEST"))
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))

	fmt.Printf("balance: %f\n", value)

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
}
