module github.com/the-medium-tech/mdl-sdk-go

go 1.14

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20220704084225-05e143d24a9e

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0
	github.com/btcsuite/btcd/btcutil v1.1.3
	github.com/golang/protobuf v1.5.0
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/hyperledger/fabric-sdk-go v1.0.0
	github.com/stretchr/testify v1.8.1
	golang.org/x/crypto v0.1.0
	golang.org/x/sys v0.5.0 // indirect
)
