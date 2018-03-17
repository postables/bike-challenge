package main 
import (
	"fmt"
	"log"
	
	"strings"
	//"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"./bike_rental"
)

const key = `{"address":"d72f0d88384c05c3d95c870ba98ac2d606939c65","crypto":{"cipher":"aes-128-ctr","ciphertext":"589a88ccbdaa312595343c907e944c8b9d9e133d443b43d4efa71c6c7cea26d0","cipherparams":{"iv":"4429d785f61dd7d37d7813a8a422d941"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f92dbdb8c2c4686a839978d9dab36601a2e950d001b6d7131dd9a22c68f32da1"},"mac":"9037da8e700215e1d79043a4fcac847768d27e28dfcd3ce16f094eb1d837f1e1"},"id":"6472fa0e-80e4-475a-8f35-ede98c37641e","version":3}`

func main() {
	// establish ipc connecton
	conn, err := ethclient.Dial("/home/solidity/.ethereum/rinkeby/geth.ipc")
	if err != nil {
		log.Fatalf("error connecting")
	}
	// authorize a conenction so we can deploy stuff and change states
	_, err = bind.NewTransactor(strings.NewReader(key), "password123")
	if err != nil {
		log.Fatalf("error unlocking account")
	}

	sale, err := bt_eth.NewBtEth(common.HexToAddress("0xd97437da366c75bf4d0508adb8d323dbb70c7cee"), conn)
	if err != nil {
		log.Fatalf("error connecting to sale contract")
	}

	owner, err := sale.Owner(nil)
	if err != nil {
		log.Fatalf("error reading owner")
	}
	fmt.Printf("Owner Address: 0x%x\n", owner)

	remainingTokens, err := sale.RemainingTokens(nil)
	if err != nil {
		log.Fatalf("error reading remaining tokens")
	}
	fmt.Printf("Remaining Tokens %v\n", remainingTokens)

	tokensSold, err := sale.NumBikeTokensSold(nil)
	if err != nil {
		log.Fatalf(" error reading tokens sold")
	}
	fmt.Printf("Tokens Sold %v\n", tokensSold)

	numberContributions, err := sale.NumContributions(nil)
	if err != nil {
		log.Fatalf("error reading number  contributions")
	}
	fmt.Printf("number of contributions %v\n", numberContributions)

	ethRaised, err := sale.EthRaised(nil)
	if err != nil {
		log.Fatalf("error reading eth raised")
	}
	fmt.Printf("eth raised %v\n", ethRaised)

	enabled, err := sale.Enabled(nil)
	if err != nil {
		log.Fatalf("error reading enabled status")
	}
	fmt.Printf("Contract enabled %v\n",  enabled)

	initialized, err := sale.Initialized(nil)
	if err != nil {
		log.Fatalf("error reading initalized")
	}
	fmt.Printf("contract initialized %v\n", initialized)

	btWei, err := sale.BtWei(nil)
	if err != nil {
		log.Fatalf("error reading btwei")
	}
	fmt.Printf("bt wei %v\n", btWei)  

	contribution, err := sale.Contributions(nil, owner)
	if err != nil {
		fmt.Printf("error reading contribution")
	}
	fmt.Printf("Contribution for 0x%x %v\n", owner, contribution)
}