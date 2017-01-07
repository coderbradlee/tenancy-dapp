# tenancy-dapp

In order to validate the contract we have created bindings for property using abigen. standalone.go file creates
a simulated blockchain and performs contract creations and transactions.

Abigen
======
`abigen --sol contracts/Property.sol --pkg main --out Property.go`

Build and Run
=============
`go build -o standalone`
`./standalone`

####Sample output

```
$ ./standalone
Owner Balance: 10000000000000
Tenant Balance: 100000000000
Rent : 50000000000
Security: 10000000000
Deduction 5000000000

Created simulated backend with government 0xa9f3c5dfa9ad81c7c56098d5ddc8bbde6b43c59c

Property deployed 0xc6b61d17a40e61284abe50683a7de6e8af1671b4
Government address setup into 0x6cc9f01123a4b4785aa5a987e248cc27b1bb75801df2e03793083418a4aa684d

Tenant balance : 100000000000
Owner balance : 9999999480026
Property balance : 0

Property validated 0xacc7249f2ae2b7b8a08c04226ff59c82e7c1c34ce4f1dd0c8945750944d69f48
Payment done  0x4a13616c116e94b1b6541fde284b4343a8066e3dac3da252a05efcaa7461057b
Tenancy terminated 0x074c466c147df1c614f242a7bf334a1316ba4f0ecdc1ddfed53498241aaa2813

Tenant balance : 44999947261
Owner balance : 10054999441724
Property balance : 0
```