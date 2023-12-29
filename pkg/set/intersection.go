package set

// IntersectionPairwise performs pairwise union on two sets
func IntersectionPairwise(set1, set2 *Set) *Set {
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
func IntersectionMultiple(sets []*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}
	if len(sets) == 1 {
		return sets[0]
	}

	mid := len(sets) / 2
	leftInter := UnionMultiple(sets[:mid])
	rightInter := UnionMultiple(sets[mid:])
	return IntersectionPairwise(leftInter, rightInter)
}

func (s *Set) Intersect(file string, o *Options) error {
	sets, err := ReadSetsFromJSON(file)
	if err != nil {
		return err
	}

	resultantSet := IntersectionMultiple(sets)

	err = writeSetToFile(resultantSet, o, "intersection")
	if err != nil {
		return err
	}

	return nil
}
