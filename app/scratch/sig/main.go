package main

import (
	"fmt"
	"log"

	"github.com/ardanlabs/blockchain/foundation/blockchain/signature"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	pkHexKey = "fae85851bdf5c9f49923722ce38f3c1defcfd3619ef5453230a58ad805499959"
	from     = "0xdd6B972ffcc631a62CAE1BB9d80b7ff429c8ebA4"
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
	pk, err := crypto.HexToECDSA(pkHexKey)
	if err != nil {
		return err
	}

	value := tx{
		From:  "0xdd6B972ffcc631a62CAE1BB9d80b7ff429c8ebA4",
		To:    "John",
		Value: 10,
	}

	v, r, s, err := signature.Sign(value, pk)
	if err != nil {
		return err
	}

	fmt.Println("V", hexutil.Encode(v.Bytes()))
	fmt.Println("R", hexutil.Encode(r.Bytes()))
	fmt.Println("S", hexutil.Encode(s.Bytes()))

	// =========================================================================

	calcAddr, err := signature.FromAddress(value, v, r, s)
	if err != nil {
		return err
	}

	addr := crypto.PubkeyToAddress(pk.PublicKey).String()
	fmt.Println("ADDR PK :", addr)
	fmt.Println("ADDR SIG:", calcAddr)

	if addr != calcAddr {
		return fmt.Errorf("NO MATCH  got %s  exp %s", addr, calcAddr)
	}

	return nil
}

func hash() error {
	v := tx{
		From:  "Bill",
		To:    "John",
		Value: 10,
	}

	hash := signature.Hash(v)
	fmt.Println("SUM 1", hash)

	v.Value = 11

	hash = signature.Hash(v)
	fmt.Println("SUM 2", hash)

	return nil
}
