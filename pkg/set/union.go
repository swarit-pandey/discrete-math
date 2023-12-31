package set

// Theorem: Associativity of Union Operation
// The union operation is associative, i.e., (A U B) U C = A U (B U C).
// This property allows us to compute the union of a collection of sets in a divide-and-conquer manner.
// We can recursively divide the collection into smaller subsets, compute their unions, and combine these results,
// ensuring the correctness of the final union irrespective of the order in which individual unions are performed.

// Union pairwise, take setA and setB, perform check take union
func unionPairwise(set1, set2 *Set) *Set {
	resultSet := NewSet()
	for element := range set1.Elements {
		resultSet.Elements[element] = true
	}

	for element := range set2.Elements {
		resultSet.Elements[element] = true
	}

	resultSet.Cardinality = len(resultSet.Elements)
	return resultSet
}

// Divide and conquer merge
func unionMultiple(sets []*Set) *Set {
	if len(sets) == 0 {
		return NewSet()
	}
	if len(sets) == 1 {
		return sets[0]
	}

	mid := len(sets) / 2
	leftUnion := unionMultiple(sets[:mid])
	rightUnion := unionMultiple(sets[mid:])
	return unionPairwise(leftUnion, rightUnion)
}

func (s *Set) Union(o *Options) error {
	sets, err := readSetsFromJSON(o)
	if err != nil {
		return err
	}

	resultantSet := unionMultiple(sets)

	err = writeSetToFile(resultantSet, o, "union")
	if err != nil {
		return err
	}

	return nil
}
