package main 
import (
	"fmt"
	"log"
	
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/common"
	"./bike_token"
)

const key = `{"address":"d72f0d88384c05c3d95c870ba98ac2d606939c65","crypto":{"cipher":"aes-128-ctr","ciphertext":"589a88ccbdaa312595343c907e944c8b9d9e133d443b43d4efa71c6c7cea26d0","cipherparams":{"iv":"4429d785f61dd7d37d7813a8a422d941"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f92dbdb8c2c4686a839978d9dab36601a2e950d001b6d7131dd9a22c68f32da1"},"mac":"9037da8e700215e1d79043a4fcac847768d27e28dfcd3ce16f094eb1d837f1e1"},"id":"6472fa0e-80e4-475a-8f35-ede98c37641e","version":3}`

func main() {
	// establish ipc connecton
	conn, err := ethclient.Dial("/home/solidity/.ethereum/rinkeby/geth.ipc")
	if err != nil {
		log.Fatalf("error connecting")
	}
	// authorize a conenction so we can deploy stuff and change states
	auth, err := bind.NewTransactor(strings.NewReader(key), "password123")
	if err != nil {
		log.Fatalf("error unlocking account")
	}
	// address, tx, token, err := ....
	address, tx, token, err := bike_token.DeployBikeToken(auth, conn)
	if err != nil {
		log.Fatalf("error ", err)
	}
	fmt.Printf("Transaction hash %x\n", tx.Hash())
	fmt.Printf("Token successfully deployed\n")
	fmt.Printf("Token Address = 0x%x\n", address)
	fmt.Printf("Sleeping for one minute\n")
	time.Sleep(250 * time.Millisecond)
	
	// using pending true will allow us to read variables while
	// the tx is pending
	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("error reading token name")
	}
	fmt.Printf("Token Name %s\n", name)
	
	symbol, err := token.Symbol(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("error reading token symbol")
	}
	fmt.Printf("Token Symbol %s\n", symbol)
	
	decimals, err := token.Decimals(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("error reading token deciamls")
	}
	fmt.Printf("Token Decimals %v\n", decimals)

	supply, err := token.TotalSupply(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("error reading token supply")
	}
	fmt.Printf("Token Supply %v\n", supply)

	owner, err := token.Owner(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("error reading token owner")
	}
	fmt.Printf("Token owner 0x%x\n", owner)

	fmt.Printf("Checking balance of contract owner\n")
	ownerBalance, err := token.BalanceOf(&bind.CallOpts{Pending: true}, owner)
	if err != nil {
		log.Fatalf("error reading balance of contract owner")
	}
	fmt.Printf("Owner balance: %v\n", ownerBalance)

}