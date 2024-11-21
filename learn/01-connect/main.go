package main

import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
)


// var infuraURL = use infura ethereum Mainnet rpc url

var anvilURL = "http://127.0.0.1:8545"

// main will connect to the Ethereum network using the Infura API and print the most recent block number.
func main() {
	client, err := ethclient.DialContext(context.Background(),anvilURL);

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
	fmt.Println(blockNumber);
}
