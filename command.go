package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"fmt"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
	"bufio"
	"flag"
	"math/big"
)

const keyFile  = "/Users/raviupreti85/Library/Ethereum/dev/keystore/" +
	"UTC--2016-12-29T21-53-55.927410990Z--429d61dc95cac25a24feffcf7db98f76d6ab3796"

func commands() {
	if len(os.Args) == 1 {
		fmt.Println("usage: <command> [<args>]")
		os.Exit(1)
	}

	balanceCommand := flag.NewFlagSet("balance", flag.ExitOnError)
	balanceAddress := balanceCommand.String("address", "", "Address of the account")

	deployCommand := flag.NewFlagSet("deploy", flag.ExitOnError)
	ejariAddress := deployCommand.String("address", "", "Address of the account")

	registryCommand := flag.NewFlagSet("registry", flag.ExitOnError)
	registryAddress := registryCommand.String("address", "", "Address of the registry")

	switch os.Args[1] {
	case "balance":
		balanceCommand.Parse(os.Args[2:])
		address := common.HexToAddress(*balanceAddress)
		getBalance(address)
	case "deploy":
		deployCommand.Parse(os.Args[2:])
		address := common.HexToAddress(*ejariAddress)
		doDeployRegistry(address)
	case "registry":
		registryCommand.Parse(os.Args[2:])
		address := common.HexToAddress(*registryAddress)
		getBalance(address)
		getRegistry(address)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

}

func getBalance(address common.Address) {
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	doGetBalance(client, address)
}

func doDeployRegistry(address common.Address) {
	file, err := os.Open(keyFile)

	transactOpts, err := bind.NewTransactor(bufio.NewReader(file), "password")

	if (err != nil) {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	deployRegistry(transactOpts, address)
}

func getRegistry(address common.Address) {
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	registry, err := NewRegistry(address, client)

	if (err != nil) {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	file, err := os.Open(keyFile)

	transactOpts, err := bind.NewTransactor(bufio.NewReader(file), "password")

	if (err != nil) {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	transaction, err := registry.IsValidTenancy(transactOpts, address, big.NewInt(10), big.NewInt(10));

	if (err != nil) {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	fmt.Printf("Registry : %v", transaction.Hash().Hex())
}

// var address = common.HexToAddress("0x68597E5Dd72CF95C77Af2b85e5670369b22AB6cA")
func doGetBalance(client *ethclient.Client, address common.Address) {
	var context = context.Background()
	balance, err := client.BalanceAt(context, address, nil)

	if (err != nil) {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	fmt.Printf("Balance : %v\n", balance)
}

/**

 */


/**
transaction, err := registry.GenerateId(transactOpts, "12 Scrubbitts Square", "WD78JR");

	if (err != nil) {
		fmt.Printf("%v", err)
		os.Exit(1);
	}

	fmt.Printf("Transaction hash : %v\n", transaction.Hash().Hex());
 */