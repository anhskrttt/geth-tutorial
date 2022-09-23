package event_log

import (
	"fmt"
	"geth-tutorial/initials"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func SubscribeToEvent() {
	fmt.Println(initials.Client)
	contractAddress := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	_ = query

	logs := make(chan types.Log)
	_ = logs
	// fmt.Println(logs)

	// sub, err := initials.Client.SubscribeFilterLogs(context.Background(), query, logs)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _ = sub

	// for {
	// 	select {
	// 	case e := <-sub.Err():
	// 		log.Fatal(e)
	// 	case vLog := <-logs:
	// 		fmt.Println(vLog) // pointer to event log
	// 	}
	// }
}
