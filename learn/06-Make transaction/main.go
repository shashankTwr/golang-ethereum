package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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


func importKeys() (key1 *keystore.Key, key2 *keystore.Key) {
	entries, err := os.ReadDir("./wallet")
	if err != nil {
		log.Fatalf("failed to read wallet directory: %v", err)
	}

	if len(entries) < 2 {
		log.Fatalf("Not enough entries")
	}

	// Use the first file found
	b1, err := os.ReadFile("./wallet/" + entries[0].Name())
	if err != nil {
		log.Fatalf("Error reading the first file: %v", err)
	}
	b2, err := os.ReadFile("./wallet/" + entries[1].Name())
	if err != nil {
		log.Fatalf("Error reading the second file: %v", err)
	}

	password := "password"
	senderKey, err := keystore.DecryptKey(b1, password)

	if err != nil {
		log.Fatalf("Error to decrypt the first file: %v", err)
	}
	receiverKey, err := keystore.DecryptKey(b2, password)
	if err != nil {
		log.Fatalf("Error to decrypt the second file: %v", err)
	}


	return senderKey, receiverKey
}


func importKsAddr() (address1 common.Address, address2 common.Address) {
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

func sendTransaction(infuraURL string, a1 common.Address, a2 common.Address) {

	client, err := ethclient.DialContext(context.Background(),infuraURL);

	if err!= nil{
		log.Fatalf("Error to create a ether client: %v", err)
	}
	defer client.Close()

	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatalf("Error to get nonce: %v", err)
	}

	// 1 ethere = 1000000000000000000 wei
	amount := big.NewInt(1000000000000000)

	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatalf("Error to get gas price: %v", err)
	}

	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)
	
	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		log.Fatalf("Error to get chain id: %v", err)
	}
	key1, _ := importKeys()
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer((chainID)), key1.PrivateKey)


	if err!= nil {
		log.Fatalf("Error to sign the transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)


	if err != nil {
		log.Fatalf("Error to send the transaction: %v", err)
	}
	
}



func main()  {
		// Import the addresses
		address1, address2 := importKsAddr()

		// Validate the addresses
		if !common.IsHexAddress(address1.Hex()) || !common.IsHexAddress(address2.Hex()) {
			log.Fatalf("Invalid Ethereum addresses: %v, %v", address1.Hex(), address2.Hex())
		}
	
		// Print the addresses
		fmt.Println("Address 1:", address1.Hex())
		fmt.Println("Address 2:", address2.Hex())
	
		// Get the Infura URL
		infuraURL, err := getEnvValue("infuraSepURL")
		if err != nil {
			log.Fatalf("Error getting Infura URL: %v", err)
		}
		fmt.Println("Infura URL:", infuraURL)
	
		// Get and print the balances
		balance1 := getbalance(infuraURL, address1)
		balance2 := getbalance(infuraURL, address2)
		fmt.Printf("Balance of Address 1: %s\n", balance1.String())
		fmt.Printf("Balance of Address 2: %s\n", balance2.String())

	
		// Send the transaction
		sendTransaction(infuraURL, address1, address2)
}