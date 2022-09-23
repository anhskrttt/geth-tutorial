package transfer

import (
	"context"
	"fmt"
	"geth-tutorial/initials"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func TransferERC20() {
	fromAddress := common.HexToAddress("0x08e1BeCDE52651095fCa66902B81Dad2057b6137")

	privateKey, err := crypto.HexToECDSA(os.Getenv("LOCAL_PRIVATE_KEY_TEST"))
	if err != nil {
		log.Fatal(err)
	}

	// ETH to send = 0
	value := big.NewInt(0)

	nonce, err := initials.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Receiver address
	toAddress := common.HexToAddress(os.Getenv("LOCAL_ADDRESS_TEST_01"))

	// Token address
	tokenAddress := common.HexToAddress("0x32B2C660eC3ca79B89ccD879bC0878F32A6D606f")

	transferFnSignature := []byte("transfer(address,uint256)")

	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("10", 10)

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	// Create data field by concatenating all data above
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// Set gas limit
	gasLimit, err := initials.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := initials.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gasLimit)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

	chainID, err := initials.Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Sign tx
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Broadcast tx
	err = initials.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
