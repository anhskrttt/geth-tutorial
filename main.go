package main

import (
	"geth-tutorial/initials"
	"geth-tutorial/transaction"
)

func init() {
	initials.LoadEnv()
	initials.SetUpClient()
}

func main() {
	// 01. Retrieve balance
	// fmt.Println(account.GetBalance(os.Getenv("LOCAL_ADDRESS_TEST")))

	// 02. Generate wallet
	// prv, pub := wallet.CreateWallet()
	// fmt.Printf("Private key: %s\nPublic key: %s\n", prv, pub)

	// 03. Generate keystore
	// keystore.GenerateKeyStore()
	// keystore.AccessAccount()

	// 04. Check if string is a valid address
	// fmt.Println(utility.IsValidAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d"))

	// fmt.Println(utility.IsContractAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"))

	// 05. Query blocks
	// fmt.Println(transaction.GetLatestBlockHeader().Number.String())
	// fmt.Println(transaction.GetFullBlock(big.NewInt(5671744)).Hash().Hex())

	// 06. Query transactions
	// blockTest := transaction.GetFullBlock(big.NewInt(5671744))
	// transaction.GetAllTransaction(blockTest)
	// transaction.GetTransactionByBlockHash()
	transaction.GetTxByTxHash()
}
