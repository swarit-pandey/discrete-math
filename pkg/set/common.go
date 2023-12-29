package set

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// TODO: Refactor this shitty code

func readSetsFromJSON(filename string) ([]*Set, error) {
	var jsonSets []JSONSet
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &jsonSets)
	if err != nil {
		return nil, err
	}

	var sets []*Set
	for _, jSet := range jsonSets {
		newSet := NewSet()
		for _, elem := range jSet.Elements {
			newSet.Elements[elem] = true
			newSet.Cardinality++
		}
		sets = append(sets, newSet)
	}

	return sets, nil
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

	var outputFilePath string
	if outputFilePath == "" {
		outputFilePath = createDefaultOutputFile()
	} else {
		outputFilePath = o.OutputFile
	}

	file, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(jsonSets)
}

func writeSetToFile(set *Set, o *Options, operation string) error {
	fmt.Println("Writing set to a file... ")
	jsonSet := make([]JSONSet, 1)

	singleSet := JSONSet{
		Name:     fmt.Sprintf("Set %s", operation),
		Elements: make([]interface{}, 0, len(set.Elements)),
	}

	for elem := range set.Elements {
		singleSet.Elements = append(singleSet.Elements, elem)
	}

	jsonSet[0] = singleSet
	var filename string
	if o.OutputFile == "" {
		filename = createDefaultOutputFile()
	} else {
		filename = o.OutputFile
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(singleSet)
}

func createDefaultOutputFile() string {
	timestamp := time.Now().Format("20060102-150405")

	return fmt.Sprintf("sets-%s.json", timestamp)
}
