package main

import (
	"context"
	"fmt"
	"os"
	"log"
	"math"
	"math/big"

	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

func getEnv(key string) (string, error) {
    value, ok := os.LookupEnv(key)
    if !ok {
        return "", fmt.Errorf("%s environment variable not set", key)
    }
    return value, nil
}



// var anvilURL = "http://127.0.0.1:8545"

// main will connect to the Ethereum network using the Infura API and print the most recent block number.
func main() {

	// Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Retrieve the environment variable
    infuraURL, err := getEnv("infuraURL")
    if err != nil {
        log.Println("Error:", err)
        return
    }

    fmt.Println("Infura URL:", infuraURL)
	client, err := ethclient.DialContext(context.Background(),infuraURL);

	if err!= nil{
		log.Fatalf("Error to create a ether client: %v", err)
	}
	defer client.Close();

	block, err := client.BlockByNumber(context.Background(),nil);
	// automatically we get the last block number

	if err != nil {
		log.Fatalf("Error to get block: %v", err)
	}
	blockNumber := block.Number().Uint64();
	fmt.Println("The latest block number is: ",blockNumber);

	addr := "0x4675C7e5BaAFBFFbca748158bEcBA61ef3b0a263"
	address := common.HexToAddress(addr);

	// context, address, blockNumber argument
	balance, err := client.BalanceAt(context.Background(), address, nil);

	if err != nil {
		log.Fatalf("Error to get balance: %v", err)
	}
	fmt.Println("balance is: ",balance);

	// 1 ether is 1e18 wei

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println("float balance is", fBalance);

	value := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)));
	fmt.Println("value is", value);
}
