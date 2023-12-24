package set

type Set struct {
	Multiset    bool                 // Multiset or not
	Cardinality int                  // Cardinality of the set
	Universe    interface{}          // Universe of the set
	Elements    map[interface{}]bool // Elements of the set
}

type Options struct {
	IgnoreDuplicate bool   // Ignore duplicate elements in the set
	SetSize         int    // Size/cardinality of the set
	NumberOfSets    int    // Number of sets to generate
	Randomize       bool   // Randomize the elements in the set
	OutputFile      string // Output file to write the sets to
	Range           int    // Starting from zero to upper bound
}

type JSONSet struct {
	Name     string        `json:"name"`
	Elements []interface{} `json:"elements"`
}

func NewSet() *Set {
	return &Set{
		Multiset:    false,
		Cardinality: 0,
		Universe:    nil,
		Elements:    make(map[interface{}]bool),
	}
}

func Sets(o *Options) {

}
