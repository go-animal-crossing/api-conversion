package source

import (
	"acnh/data/samples"
	"log"
	"testing"
)

func Test_Simple_Conversion(t *testing.T) {
	s := Source{}

	// setup mock loaded data on the source item
	all := make([]Item, 0)
	fishJSON := []byte(samples.MultipleFish())
	all, _ = s.unmarshal(fishJSON)

	// add fish config to the data
	s.AddConfigToData("fish", &all)

	s.All = all
	// run the convert to get target type and errors
	tg, e := s.Convert()

	if e != nil {
		log.Fatalf("Error during conversion:\n%v\n", e)
	}

	found := tg.All
	ok := 0
	for _, r := range found {
		if r.Type.Slug == "fish" {
			ok++
		}
	}

	if ok == 0 {
		log.Fatalf("Converted items did not contain 'fish'\n%v\n", found)
	}

	if ok != 2 {
		log.Fatalf("Converted fish incorrect length: %v", found)
	}

}
