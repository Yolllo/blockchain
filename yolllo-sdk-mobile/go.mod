module yolllo-sdk-mobile

go 1.17

require github.com/ElrondNetwork/elrond-sdk-erdgo v1.1.0

require (
	github.com/ElrondNetwork/elrond-go-core v1.1.15 // indirect
	github.com/ElrondNetwork/elrond-go-crypto v1.0.1 // indirect
	github.com/ElrondNetwork/elrond-go-logger v1.0.7 // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce // indirect
	github.com/denisbrodbeck/machineid v1.0.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/mobile v0.0.0-20220722155234-aaac322e2105 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/ElrondNetwork/elrond-sdk-erdgo v1.1.0 => ../yolllo-sdk-go

replace github.com/ElrondNetwork/arwen-wasm-vm/v1_2 v1.2.40 => github.com/ElrondNetwork/wasm-vm v1.2.40

replace github.com/ElrondNetwork/arwen-wasm-vm/v1_3 v1.3.40 => github.com/ElrondNetwork/wasm-vm v1.3.40

replace github.com/ElrondNetwork/arwen-wasm-vm/v1_4 v1.4.54-rc3 => github.com/ElrondNetwork/wasm-vm v1.4.54-rc3

replace github.com/ElrondNetwork/elrond-go v1.2.32 => ../yolllo-go

replace github.com/ElrondNetwork/elrond-go-core v1.1.15 => ../yolllo-go-core

replace github.com/ElrondNetwork/elrond-deploy-go v0.0.0-20211109071733-74552b78d0a2 => ../yolllo-deploy-go
