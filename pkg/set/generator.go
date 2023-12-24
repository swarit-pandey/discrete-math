package set

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
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

func writeSetsToFile(sets []*Set, o *Options) error {
	fmt.Println("Writing sets to file...")

	jsonSets := make([]JSONSet, len(sets))

	for i, set := range sets {
		jsonSet := JSONSet{
			Name:     fmt.Sprintf("Set %d", i+1),
			Elements: make([]interface{}, 0, len(set.Elements)),
		}

		for elem := range set.Elements {
			jsonSet.Elements = append(jsonSet.Elements, elem)
		}

		jsonSets[i] = jsonSet
	}

	file, err := os.Create(o.OutputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(jsonSets)
}
