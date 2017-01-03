package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func deployRegistry(transactOpts *bind.TransactOpts) *Registry {
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		fmt.Printf("%v", err)
	}

	address, transaction, registry, err := DeployRegistry(transactOpts, client)

	if (err != nil) {
		fmt.Printf("Error : %v", err)
	}

	fmt.Printf("Transaction hash: %v", transaction.Hash().Hex());
	fmt.Printf("Registry Address: %v", address.Hex());

	return registry

}

func deployProperty(transactOpts *bind.TransactOpts, street string, postcode string) *Property {
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		fmt.Printf("%v", err)
	}

	address, transaction, property, err := DeployProperty(transactOpts, client, street, postcode)

	if (err != nil) {
		fmt.Printf("Error : %v", err)
	}

	fmt.Printf("Transaction hash: %v", transaction.Hash().Hex());
	fmt.Printf("Registry Address: %v", address.Hex());

	return property
}

