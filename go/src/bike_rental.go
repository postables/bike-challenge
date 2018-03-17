package main 
import (
	"fmt"
	"log"
	"math/big"

	"strings"
	"time"

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
	auth, err := bind.NewTransactor(strings.NewReader(key), "password123")
	if err != nil {
		log.Fatalf("error unlocking account")
	}

	rental, err := bike_rental.NewBikeRental(common.HexToAddress("0x78d85767288fa002fa584336e35458bb4beb61e0"), conn)
	if err != nil {
		log.Fatalf("error connecting to sale contract")
	}

	count, err := rental.BikeCount(nil)
	if err != nil {
		log.Fatalf("error reading bike count")
	}
	fmt.Printf("Bike Count %v\n", count)

	if count == big.NewInt(0) {
		fmt.Printf("Adding bike\n")
		tx, err := rental.AddBike(auth, big.NewInt(100), big.NewInt(500000))
		if err != nil {
			log.Fatalf("error adding bike, try increasing the bike identifier")
		}
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
	} else {
		fmt.Printf("bike already added so we arent adding one \n")
	}

	bike, err := rental.Bikes(nil, big.NewInt(1))
	if err != nil {
		log.Fatalf("error retrieving bike struct")
	}
	fmt.Printf("Bike Struct %v\n", bike)
}