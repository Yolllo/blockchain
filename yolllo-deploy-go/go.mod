module github.com/ElrondNetwork/elrond-deploy-go

go 1.13

require (
	github.com/ElrondNetwork/elrond-go v1.2.32
	github.com/ElrondNetwork/elrond-go-core v1.1.15
	github.com/ElrondNetwork/elrond-go-crypto v1.0.1
	github.com/ElrondNetwork/elrond-go-logger v1.0.7
	github.com/ElrondNetwork/elrond-vm-common v1.3.12
	github.com/stretchr/testify v1.7.1
	github.com/urfave/cli v1.22.9
)

replace github.com/ElrondNetwork/arwen-wasm-vm/v1_2 v1.2.40 => github.com/ElrondNetwork/wasm-vm v1.2.40

replace github.com/ElrondNetwork/arwen-wasm-vm/v1_3 v1.3.40 => github.com/ElrondNetwork/wasm-vm v1.3.40

replace github.com/ElrondNetwork/arwen-wasm-vm/v1_4 v1.4.54-rc3 => github.com/ElrondNetwork/wasm-vm v1.4.54-rc3

replace github.com/ElrondNetwork/elrond-go v1.2.32 => ../yolllo-go

replace github.com/ElrondNetwork/elrond-go-core v1.1.15 => ../yolllo-go-core
