package keystore

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func GenerateKeyStore() {
	ks := keystore.NewKeyStore("./keys", keystore.StandardScryptN, keystore.StandardScryptP)
	password := os.Getenv("KEYSTORE_PASSWORD")
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

func AccessAccount() {
	file := "./keys/UTC--2022-09-22T03-09-33.546135689Z--8ed3c142b7bf9d0ff82b0cf79cf544e8bc9d9a80"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := os.Getenv("KEYSTORE_PASSWORD")
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
