package source

import (
	"acnh/data/samples"
	"log"
	"testing"
)

func Test_Source_Unmarshal(t *testing.T) {
	var e error
	var data []Item
	s := Source{}
	// mock data loading
	fishJSON := []byte(samples.MultipleFish())
	data, e = s.unmarshal(fishJSON)
	e = s.AddConfigToData("fish", &data)

	if e != nil {
		log.Fatalf("Failed to unmarshal before converting:\n%v\n", e)
	}

	if len(data) != 2 {
		log.Fatalf("Failed to load 2 fish:\n %v\n", data)
	}
	// assign
	s.All = data
}

func Test_Config_Data_Adding(t *testing.T) {

	var data []Item
	s := Source{}
	// fish should be found in the Config data struct, so no error
	e := s.AddConfigToData("fish", &data)
	if e != nil {
		log.Fatalf("Failed to find fish in the config data source: %v\n", e)
	}
	// foo bar does not exist, so should throw an error
	e = s.AddConfigToData("foobar", &data)
	if e == nil {
		log.Fatalln("Failed to throw an error for foobar config")
	}
	// test dummy data gets the config for fish added
	fishJSON := []byte(samples.MultipleFish())
	data, e = s.unmarshal(fishJSON)
	e = s.AddConfigToData("fish", &data)

	cfg := data[0].Config
	if cfg.Name != "fish" {
		log.Fatalln("Failed to correctly attach fish config data")
	}

}
