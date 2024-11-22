package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)
func getEnv(key string) (string, error) {
    value, ok := os.LookupEnv(key)
    if !ok {
        return "", fmt.Errorf("%s environment variable not set", key)
    }
    return value, nil
}
func createKs() {
    ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
    password := "password"
    account, err := ks.NewAccount(password)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(account.Address.Hex()) 
}



// importKs reads the first two keystore files from the `./wallet` directory,
// decrypts them using a predefined password, and returns the corresponding ``
// Ethereum public addresses. It expects at least two keystore files to exist 
// in the directory and logs a fatal error if any operation fails.
func importKs() (address1 common.Address, address2 common.Address) {
	// Read files from the `./wallet` directory
	entries, err := os.ReadDir("./wallet")
	if err != nil {
		log.Fatalf("failed to read wallet directory: %v", err)
	}

	if len(entries) < 2 {
		log.Fatalf("Not enough entries")
	}

	// Use the first file found
	b, err := os.ReadFile("./wallet/" + entries[0].Name())

	if err != nil {
		log.Fatalf("Error to read the file: %v", err)
	}
	password := "password"
	key, err := keystore.DecryptKey(b, password)

	if err != nil {
		log.Fatalf("Error to decrypt the file: %v", err)
	}

	address1 = crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	// Use the second file found
	b1, err := os.ReadFile("./wallet/" + entries[1].Name())

	if err != nil {
		log.Fatalf("Error to read the file: %v", err)
	}
	key1, err := keystore.DecryptKey(b1, password)

	if err != nil {
		log.Fatalf("Error to decrypt the file: %v", err)
	}

	address2 = crypto.PubkeyToAddress(key1.PrivateKey.PublicKey)

	return address1, address2
}

func getEnvValue(key string) (string, error) {

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	value, err:= getEnv(key)
	if err!=nil {
		return "", fmt.Errorf("%s environment variable not set", key)
	}
	return value, nil
}

// getbalance connects to the Ethereum network using the provided Infura URL and
// prints the balance of the provided address.
//
// If the connection to the network fails, it logs the error and stops executing.
// If the balance retrieval fails, it logs the error and stops executing.
//
func getbalance(infuraURL string, address common.Address) (valance *big.Int) {
	client, err := ethclient.DialContext(context.Background(),infuraURL);

	if err!= nil{
		log.Fatalf("Error to create a ether client: %v", err)
	}
	defer client.Close();

	balance, err := client.BalanceAt(context.Background(),address, nil)
	if err != nil {
		log.Fatalf("Error to get balance: %v", err)
	}
	return balance
}

func main()  {
	//createKs()

	// get the address of the two accounts
	address1, address2 := importKs()

	// print the two addresses
	fmt.Println("address 1 is ",address1.Hex())
	fmt.Println("address 2 is ",address2.Hex())

	// get the infura url
	infuraURL, err := getEnvValue("infuraSepURL")
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Println("Infura URL:", infuraURL)
	
	// get the balance of the two addresses
	balance1 := getbalance(infuraURL, address1)
	balance2 := getbalance(infuraURL, address2)

	// print the balances
	fmt.Println("balance of address 1 is ",balance1)
	fmt.Println("balance of address 2 is ",balance2)

}