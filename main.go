package main

import (
	"fmt"
	"geth-tutorial/initials"
	"geth-tutorial/utility"
)

func init() {
	initials.LoadEnv()
	initials.SetUpClient()
}

func main() {
	// fmt.Println(account.GetBalance(os.Getenv("LOCAL_ADDRESS_TEST")))

	// prv, pub := wallet.CreateWallet()
	// fmt.Printf("Private key: %s\nPublic key: %s\n", prv, pub)

	// keystore.GenerateKeyStore()
	// keystore.AccessAccount()

	// fmt.Println(utility.IsValidAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d"))

	fmt.Println(utility.IsContractAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"))
}
