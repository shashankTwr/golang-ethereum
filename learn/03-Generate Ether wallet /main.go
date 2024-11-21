package main

import (
	"log"
	"fmt"


	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

)

func main() {

	// 
	pvk,err := crypto.GenerateKey();

	if err!= nil{
		log.Fatalf("Error to create a private key  %v", err)
	}

	pData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(pData))

	puData := crypto.FromECDSAPub(&pvk.PublicKey) // passing as an address
	fmt.Println(hexutil.Encode(puData))// publicKey generated

	// address
	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
