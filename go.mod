module github.com/the-medium-tech/mdl-sdk-go

go 1.14

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20220704084225-05e143d24a9e

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0
	github.com/btcsuite/btcd/btcutil v1.1.3
	github.com/hyperledger/fabric-sdk-go v1.0.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
)
