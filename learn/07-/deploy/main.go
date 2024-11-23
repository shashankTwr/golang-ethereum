package main

import (
	"context"
	"fmt"
	"log"
	"os"

	todo "github.com/shashanktwr/todo"

	"github.com/ethereum/go-ethereum/accounts/keystore"
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

func main(){
	key1, _ := importKeys()

	key, err := keystore.DecryptKey(key1, "password")

	if err!= nil {
		log.Fatal(err)
	}

	infuraSepURL, err := getEnvValue("infuraSepURL")
	if err != nil {
		log.Fatalf("Error getting Infura URL: %v", err)
	}
	fmt.Println("Infura URL:", infuraSepURL)

	client, err := ethclient.Dial(infuraSepURL)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	// get nonce for the account, gasPrice suggested, chainId of the network
	nonce, err := client.PendingNonceAt(context.Background(), add)

	if err!= nil {
		log.Fatal(err)
	}


	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err!= nil {
		log.Fatal(err)
	}

	chainID, err = client.NetworkID(context.Background())
	if err!= nil {
		log.Fatal(err)
	}


	// create a transaction
	todo.DeployTodo()

}