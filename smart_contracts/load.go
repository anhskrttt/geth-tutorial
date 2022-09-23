package smart_contracts

import (
	"errors"
	"fmt"
	"geth-tutorial/initials"
	"geth-tutorial/store"
	"geth-tutorial/utility"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func GetContract() (*store.Store, error) {
	// instance := &store.Store{}
	// var err error
	addrString := "0x686ed89842b801D83b570b90997AB4f38b50Ba1a"

	// Check if address is a valid contract address
	if !utility.IsContractAddress(addrString) {
		fmt.Println("Invalid address")
		return &store.Store{}, errors.New("invalid contract address")
	}

	address := common.HexToAddress(addrString)

	instance, err := store.NewStore(address, initials.Client)
	if err != nil {
		log.Fatal(err)
	}

	// _ = instance // we'll be using this in the next section
	return instance, nil
}
