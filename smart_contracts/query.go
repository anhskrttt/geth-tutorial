package smart_contracts

import (
	"fmt"
	"log"
)

func QueryContract() {
	instance, err := GetContract()
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}

func QueryItem() {
	instance, err := GetContract()
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	copy(key[:], []byte("foo"))

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // "bar"
}
