package smart_contracts

import (
	"context"
	"fmt"
	"geth-tutorial/initials"
	"geth-tutorial/store"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func DeployContract() {
	fromAddress := common.HexToAddress("0x08e1BeCDE52651095fCa66902B81Dad2057b6137")

	privateKey, err := crypto.HexToECDSA(os.Getenv("LOCAL_PRIVATE_KEY_TEST"))
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := initials.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// gasLimit, err := initials.Client.EstimateGas(context.Background(), ethereum.CallMsg{
	// 	From: fromAddress,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	gasPrice, err := initials.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)    // in wei
	auth.GasLimit = uint64(4712388)// in units
	// auth.GasLimit = gasLimit // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, initials.Client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	_ = instance
}
