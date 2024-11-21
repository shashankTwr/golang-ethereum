package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {

	// // Create a new keystore
	// key := keystore.NewKeyStore("./wallet",keystore.StandardScryptN,keystore.StandardScryptP)
	// // password string phrase
	password := "password"

	// // new account method of keystore takes a string and returns an account and an error
	// acct, err :=  key.NewAccount(password)

	// if err!= nil{
	// 	log.Fatalf("Error to create a account: %v", err)

	// }
	// fmt.Println(acct.Address)

	// decrypt the generated key store file
	
	b, err := os.ReadFile("./wallet/UTC--2024-11-21T12-08-55.862126328Z--955ac1e3bb8c8bb09997a9142a4bdee6a6a0630d")

	if err != nil {
		log.Fatalf("Error to read the file: %v", err)
	}

	key, err := keystore.DecryptKey(b, password)

	if err != nil {
		log.Fatalf("Error to decrypt the file: %v", err)
	}
	
	pData := crypto.FromECDSA(key.PrivateKey)
	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey) // passing as an address
	fmt.Println("Private Key: ",hexutil.Encode(pData))
	fmt.Println("Public Key:",hexutil.Encode(pubData))
	fmt.Println("Public address:",crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}