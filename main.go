package main

import (
	"context"
	"fmt"
	"geth-tutorial/initials"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {
	initials.LoadEnv()
	// initials.SetUpClient()
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
	// transaction.GetTxByTxHash()
	// transfer.TransferETH()
	// transfer.TransferERC20()

	// smart_contracts.DeployContract()

	// smart_contracts.QueryContract()

	// smart_contracts.WriteContract()

	// smart_contracts.QueryItem()

	// smart_contracts.QueryERC20()

	//event_log.SubscribeToEvent()

	// Using Websocket
	client, err := ethclient.Dial(os.Getenv("INFURA_SOCKET_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// USDT contract
	contractAddress := common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
