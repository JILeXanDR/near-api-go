package nearapi

import (
	"math/big"
)

type Balance struct {
	original string
	value    *big.Float
}

// String returns original yoctoNEAR amount.
func (b Balance) String() string {
	return b.original
}

// ToNEARs returns NEAR amount.
func (b Balance) ToNEARs() string {
	bi := &big.Float{}
	bi.Quo(b.value, oneNEAR.value)

	return bi.String()
}

func (b Balance) ToNEARsF64() (float64, error) {
	res := new(big.Float).Quo(b.value, oneNEAR.value)

	f64, _ := res.Float64()

	return f64, nil
}
