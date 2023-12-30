package set

// Need to handle the case when there are more
// than just two sets. Not adding it in the cmd
// for now.

func SetDiff(set1, set2 *Set) *Set {
	differenceSet := NewSet()

	for element := range set1.Elements {
		if !set2.Elements[element] {
			differenceSet.Elements[element] = true
		}
	}

	return differenceSet
}

func SymmetricDiff(set1, set2 *Set) *Set {
	return unionPairwise(SetDiff(set1, set2), SetDiff(set2, set1))
}
