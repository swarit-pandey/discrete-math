package combinatorics

import (
	"errors"
	"math/big"
)

func (c *CombinatorialObject) CalculatePermutations(o *Option) error {
	var n uint
	if o.InputFile == "" {
		n = uint(len(o.Elements))
	} else {
		n = uint(len(c.InputSet.Elements))
	}

	k := o.K
	if k < 1 {
		return errors.New("value of K has to be a positive non-zero integer")
	}

	var err error
	if o.Repetition {
		c.ResultIntegral, err = PermuteWithReps(n, k)
	} else {
		c.ResultIntegral, err = PermuteWithoutReps(n, k)
	}

	if err != nil {
		return err
	}
	return nil
}

func PermuteWithoutReps(n, k uint) (*big.Int, error) {
	nFact, err := Factorial(int(n))
	if err != nil {
		return nil, err
	}

	nkFact, err := Factorial(int(n - k))
	if err != nil {
		return nil, err
	}

	return new(big.Int).Div(nFact, nkFact), nil
}

func PermuteWithReps(n, k uint) (*big.Int, error) {
	result := big.NewInt(1)
	for i := uint(0); i < k; i++ {
		result.Mul(result, big.NewInt(int64(k)))
	}
	return result, nil
}
