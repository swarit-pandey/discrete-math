package combinatorics

import (
	"math/big"

	"github.com/swarit-pandey/discrete-math/pkg/set"
)

type Option struct {
	InputFile  string // Input file
	OutputFile string // Output file
	Elements   []int  // User defined element set
	Depth      uint   // Depth of the tree when needed
	Repetition bool   // Is rep allowed or not
	Circular   bool   // Used in case to find circular permutations
	Multiset   bool   // Multiset permutation and combination
	Dearange   bool   // Dearange elements
	Number     uint   // Generic number whenever needed
	N          uint   // N in 'N c K'
	K          uint   // K in 'N c K'
	R          uint   // R as an extra paramter when needed
}

type CombinatorialObject struct {
	InputSet       set.JSONSet
	ResultantSets  []set.JSONSet
	ResultIntegral *big.Int
	CurrentIndex   uint
}

func Object() *CombinatorialObject {
	return &CombinatorialObject{
		InputSet:       set.JSONSet{},
		ResultantSets:  nil,
		ResultIntegral: big.NewInt(0),
		CurrentIndex:   0,
	}
}
