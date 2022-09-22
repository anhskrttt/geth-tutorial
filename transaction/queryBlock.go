package transaction

import (
	"context"
	"geth-tutorial/initials"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

func GetLatestBlockHeader() *types.Header {
	header, err := initials.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return header
}

func GetFullBlock(blockNumber *big.Int) *types.Block {
	// blockNumber := big.NewInt(5671744)
	block, err := initials.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(block.Number().Uint64())     // 5671744
	// fmt.Println(block.Time())                // 1527211625
	// fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	// fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	// fmt.Println(len(block.Transactions()))   // 144

	return block
}

func CountAllTxInBlock(block *types.Block) uint {
	count, err := initials.Client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	return count

}
