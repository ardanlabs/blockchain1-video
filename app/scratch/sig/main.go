package main

import (
	"fmt"
	"log"

	"github.com/ardanlabs/blockchain/foundation/blockchain/signature"
)

type tx struct {
	From  string
	To    string
	Value int
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	v := tx{
		From:  "Bill",
		To:    "John",
		Value: 10,
	}

	hash := signature.Hash(v)
	fmt.Println(hash)

	v.Value = 11

	hash = signature.Hash(v)
	fmt.Println(hash)

	return nil
}
