package helper

import (
	"fmt"
	"math/big"
)

func BigintToHex(value *big.Int) string {
	valueHex := fmt.Sprintf("%x", value)
	if len(valueHex)%2 != 0 {
		valueHex = "0" + valueHex
	}
	return valueHex
}

func Int64ToHex(value int64) string {
	valueHex := fmt.Sprintf("%x", value)
	if len(valueHex)%2 != 0 {
		valueHex = "0" + valueHex
	}
	return valueHex
}
