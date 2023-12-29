package set

import "fmt"

func generatePowerset(set *Set) []*Set {
	powersetSize := 1 << len(set.Elements) // 2^n subsets
	powerset := make([]*Set, 0, powersetSize)

	elements := make([]interface{}, 0, len(set.Elements))
	for elem := range set.Elements {
		elements = append(elements, elem)
	}

	for counter := 0; counter < powersetSize; counter++ {
		subset := NewSet()
		for j, elem := range elements {
			if counter&(1<<j) != 0 {
				subset.Elements[elem] = true
			}
		}
		powerset = append(powerset, subset)
	}

	return powerset
}

func (s *Set) Powerset(filepath string, o *Options) error {
	fmt.Println("Generating subsets...")

	sets, err := readSetsFromJSON(filepath)
	if err != nil {
		return err
	}

	var powersets [][]*Set
	for _, set := range sets {
		powersets = append(powersets, generatePowerset(set))
	}

	for _, powerset := range powersets {
		err = writeSetsToFile(powerset, o)
		if err != nil {
			return err
		}
	}

	return nil
}
