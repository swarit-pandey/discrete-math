package set

import (
	"fmt"
	"math/rand"
	"time"
)

func (s *Set) Generate(o *Options) ([]*Set, error) {
	fmt.Println("Generating sets...")

	sets := generateSets(o)
	err := writeSetsToFile(sets, o)
	if err != nil {
		return nil, err
	}
	return sets, nil
}

func generateSets(o *Options) []*Set {
	noOfSets := o.NumberOfSets
	sets := make([]*Set, noOfSets)

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
    
	for i := 0; i < noOfSets; i++ {
		sets[i] = NewSet()
		for sets[i].Cardinality < o.SetSize {
			random := rng.Intn(o.Range)
			if !o.IgnoreDuplicate || !sets[i].Elements[random] {
				sets[i].Elements[random] = true
				sets[i].Cardinality++
			}
		}
	}

	return sets
}
