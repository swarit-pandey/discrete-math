package combinatorics

import (
	"errors"
	"math/big"
)

func Factorial(n int) (*big.Int, error) {
	if n > 20 {
		return nil, errors.New("number too big to calculate the factorial")
	}

	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}

	return result, nil
}
