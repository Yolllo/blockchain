package core

// AddressHandler will handle different implementations of an address
type AddressHandler interface {
	AddressAsBech32String() string
	AddressBytes() []byte
	AddressSlice() [32]byte
	IsValid() bool
	IsInterfaceNil() bool
}
