package source

import (
	"acnh/data/samples"
	"encoding/json"
	"log"
	"testing"
)

func Test_Single_Item_From_JSON(t *testing.T) {

	fishJSON := []byte(samples.SingleFish())

	fish := Item{}
	err := json.Unmarshal(fishJSON, &fish)

	if err != nil {
		log.Fatalf("Failed to unmarshall to fish model:\n%v\n", err)
	}
	// compare things
	if fish.ID != 1 {
		log.Fatalf("Fish ID does not match\n")
	}
	if fish.Price != 900 {
		log.Fatalf("Fish price does not match\n")
	}
	if fish.PriceCj != 1350 {
		log.Fatalf("Fish price-cj does not match\n")
	}

	if fish.ImageURI != "https://acnhapi.com/v1/images/fish/1" {
		log.Fatalf("Fish image URI does not match\n")
	}

	if fish.Names.EuEn != "bitterling" {
		log.Fatalf("Fish Name does not match\n")
	}

	if fish.Availability.IsAllDay != true {
		log.Fatalf("Fish availability all day does not match\n")
	}
	if fish.Availability.IsAllYear != false {
		log.Fatalf("Fish all year does not match\n")
	}

	if len(fish.Availability.TimeArray) < 1 {
		log.Fatalf("Fish availability time array length incorrect\n")
	}
	if fish.Availability.TimeArray[0] != 0 {
		log.Fatalf("Fish availability time array first entry doesnt match\n")
	}

}

func Test_Multiple_Item_From_JSON(t *testing.T) {

	fishJSON := []byte(samples.MultipleFish())

	fishes := make([]Item, 0)
	err := json.Unmarshal(fishJSON, &fishes)

	if err != nil {
		log.Fatalf("Failed to unmarshall array of fish\n")
	}

	if len(fishes) != 2 {
		log.Fatalf("Incorrect length after unmarshall\n")
	}

}

func Test_Complex_Items_From_JSON(t *testing.T) {

	itemsJSON := []byte(samples.ComplexItems())
	items := make([]Item, 0)

	err := json.Unmarshal(itemsJSON, &items)

	if err != nil {
		log.Fatalf("Failed to unmarshall array of items\n")
	}

	if len(items) != 6 {
		log.Fatalf("Incorrect length after unmarshall\n")
	}

}
