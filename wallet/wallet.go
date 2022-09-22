package wallet

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func CreateWallet() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	// Hexa string
	// Used for signing tx
	// fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	// Public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// Address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// fmt.Println(address)
	return hexutil.Encode(privateKeyBytes), address
}
