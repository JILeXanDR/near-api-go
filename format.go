package nearapi

import (
	"math/big"
)

var oneNEAR = ParseAmount("1000000000000000000000000")

// ParseAmount parses yoctoNEAR amount.
func ParseAmount(str string) Balance {
	bi := &big.Float{}
	bi.SetString(str)

	return Balance{str, bi}
}
