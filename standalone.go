package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"time"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// initialize ethers that will be used
	var govBal = big.NewInt(1000000000000)
	var ownBal = big.NewInt(10000000000000)
	var tenBal = big.NewInt(100000000000)

	var rent = big.NewInt(50000000000) // half of tenant's balance
	var security = big.NewInt(10000000000) // one tenth of tenant's balance
	var deduction = big.NewInt(5000000000) // half of security

	fmt.Printf("Owner Balance: %v\nTenant Balance: %v\nRent : %v\nSecurity: %v\nDeduction %v\n", ownBal, tenBal, rent, security, deduction)

	governmentKey, _ := crypto.GenerateKey()
	governmentAuth := bind.NewKeyedTransactor(governmentKey)
	government := governmentAuth.From
	governmentAccount := core.GenesisAccount{Address: government, Balance: govBal}

	ownerKey, _ := crypto.GenerateKey()
	ownerAuth := bind.NewKeyedTransactor(ownerKey)
	ownerAccount := core.GenesisAccount{Address: ownerAuth.From, Balance: ownBal}

	tenantKey, _ := crypto.GenerateKey();
	tenantAuth := bind.NewKeyedTransactor(tenantKey)
	tenantAccount := core.GenesisAccount{Address: tenantAuth.From, Balance: tenBal}

	fmt.Printf("Created simulated backend with government %v\n", governmentAuth.From.Hex())
	backend := backends.NewSimulatedBackend(governmentAccount, ownerAccount, tenantAccount)


	// owner clicks on the map and create's a property
	propertyAddress, _, property, err := DeployProperty(ownerAuth, backend, "0", "0", rent, security)

	if (err != nil) {
		os.Exit(1)
	}

	backend.Commit()
	fmt.Printf("Property deployed %v\n", propertyAddress.Hex())

	// setup government address into property
	transaction, err := property.SetGovernment(ownerAuth, government);
	if (err != nil) {
		os.Exit(1)
	}
	backend.Commit()
	fmt.Printf("Government address setup into %v\n", transaction.Hash().Hex());


	// validate property
	transaction, err = property.Validate(governmentAuth)

	if (err != nil) {
		os.Exit(1)
	}
	backend.Commit()
	showBalances(backend, ownerAuth.From, tenantAuth.From, propertyAddress)

	fmt.Printf("Property validated %v\n", transaction.Hash().Hex());
	

	tenantAuth.Value = big.NewInt(0).Add(rent, security)
	// tenant and owner are both happy at this moment so payment is done
	transaction, err = property.Pay(tenantAuth, big.NewInt(time.Now().Unix()), big.NewInt(time.Now().Unix()))
	if (err != nil) {
		os.Exit(1)
	}
	backend.Commit()
	fmt.Printf("Payment done by %v\n", transaction.Hash().Hex())

	transaction, err = property.Terminate(ownerAuth, deduction)
	backend.Commit()
	fmt.Printf("Tenancy terminated %v\n", transaction.Hash().Hex())

	showBalances(backend, ownerAuth.From, tenantAuth.From, propertyAddress)
}

func showBalances(backend *backends.SimulatedBackend, owner common.Address, tenant common.Address, property common.Address)  {
	tenantBalance, _ := backend.BalanceAt(nil, tenant, nil)
	fmt.Printf("Tenant balance : %v\n", tenantBalance)

	ownerBalance, _ := backend.BalanceAt(nil, owner, nil)
	fmt.Printf("Owner balance : %v\n", ownerBalance)

	propertyBalance, _ := backend.BalanceAt(nil, property, nil)
	fmt.Printf("Property balance : %v\n", propertyBalance)
}
