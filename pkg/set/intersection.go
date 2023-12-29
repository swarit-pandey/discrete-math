package set

// IntersectionPairwise performs pairwise union on two sets
func intersectionPairwise(set1, set2 *Set) *Set {
	resultSet := NewSet()

	// Go through the smaller set first
	if len(set1.Elements) > len(set2.Elements) {
		set1, set2 = set2, set1
	}

	for element := range set1.Elements {
		if set2.Elements[element] {
			resultSet.Elements[element] = true
		}
	}

	resultSet.Cardinality = len(resultSet.Elements)
	return resultSet
}

// IntersetctionMultiple will divide and conquer the sets and performs union pairwise
func intersectionMultiple(sets []*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}
	if len(sets) == 1 {
		return sets[0]
	}

	mid := len(sets) / 2
	leftInter := intersectionMultiple(sets[:mid])
	rightInter := intersectionMultiple(sets[mid:])
	return intersectionPairwise(leftInter, rightInter)
}

func (s *Set) Intersect(file string, o *Options) error {
	sets, err := readSetsFromJSON(file)
	if err != nil {
		return err
	}

	resultantSet := intersectionMultiple(sets)

	err = writeSetToFile(resultantSet, o, "intersection")
	if err != nil {
		return err
	}

	return nil
}
