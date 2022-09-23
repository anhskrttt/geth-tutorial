package initials

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client

func SetUpClient() {
	var err error
	// Set up client
	// Source: Public RPC Node (https://docs.bscscan.com/misc-tools-and-utilities/public-rpc-nodes)
	// Chainlist: https://chainlist.org/
	// client, err := ethclient.Dial(os.Getenv("BSC_RPC_NODE"))
	Client, err = ethclient.Dial(os.Getenv("LOCAL_NODE"))
	// Client, err = ethclient.Dial(os.Getenv("ETH_RPC_NODE"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = Client // we'll use this in the upcoming sections
}
