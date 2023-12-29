package set

func getComp(set, universalSet *Set) *Set {
	complementSet := NewSet()

	for element := range universalSet.Elements {
		if !set.Elements[element] {
			complementSet.Elements[element] = true
		}
	}

	complementSet.Cardinality = len(complementSet.Elements)
	return complementSet
}

// Assumption is set2 is universal set here
func (s *Set) Complement(set1, set2 *Set, o *Options) error {
	resultSet := getComp(set1, set2)

	err := writeSetToFile(resultSet, o, "complement")
	if err != nil {
		return err
	}

	return nil
}
