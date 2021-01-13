package source

import (
	"log"
	"strings"
	"testing"
)

func Test_ConvertID(t *testing.T) {
	i := sampleSimpleSourceItem()
	id := i.convertID()

	if id != "1" {
		log.Fatalln("Failed to convert simple id to string")
	}

}

func Test_ConvertName(t *testing.T) {
	i := sampleSimpleSourceItem()
	name := i.convertName()

	if name != strings.Title(i.Names.EuEn) {
		log.Fatalf("Failed to convert name: [%v]\n", name)
	}
}

func Test_ConvertType(t *testing.T) {
	i := sampleSimpleSourceItem()

	converted := i.convertType()

	if converted.Title != "Fish" {
		log.Fatalf("Failed to convert type data: \n%v\n", converted)
	}
}

func Test_ConvertPrices(t *testing.T) {
	i := sampleSimpleSourceItem()
	p := i.convertPrices()

	if p.Store != i.Price {
		log.Fatalf("Failed to convert store price data: [%v]\n", p.Store)
	}

	if p.Flick != i.PriceFlick {
		log.Fatalf("Failed to convert flick price data: [%v]\n", p.Flick)
	}

	if p.Cj != 0 {
		log.Fatalf("Failed to convert empty data: [%v]\n", p.Cj)
	}
}

func Test_ConvertImages(t *testing.T) {
	i := sampleSimpleSourceItem()
	img := i.convertImages()

	if img.MainSource != i.ImageURI {
		log.Fatalf("Failed to convert image uri: [%v]\n", img.Main)
	}

	if img.ThumbSource != i.IconURI {
		log.Fatalf("Failed to convert icon uri: [%v]\n", img.Thumb)
	}
}

func Test_ConvertAttributes(t *testing.T) {
	i := sampleSimpleSourceItem()
	attr := i.convertAttributes()

	if attr.Shadow != i.Shadow {
		log.Fatalf("Failed to convert shadow: [%v]\n", attr.Shadow)
	}
	if attr.Speed != i.Speed {
		log.Fatalf("Failed to convert speed: [%v]\n", attr.Speed)
	}

}

func Test_ConvertAvailabilty(t *testing.T) {
	i := sampleSimpleSourceItem()
	av := i.convertAvailability()
	// only testing the simple top level items here, further tests will check
	// the lower items
	if av.Location != i.Availability.Location {
		log.Fatalf("Failed to convert location: [%v]\n", av.Location)
	}
	if av.Rarity != i.Availability.Rarity {
		log.Fatalf("Failed to convert Rarity: [%v]\n", av.Rarity)
	}
}

func Test_ConvertTimes(t *testing.T) {
	i := sampleSimpleSourceItem()
	times := i.convertTimes()
	// testing time field conversion as a sub set of the availability
	if times.Always != i.Availability.IsAllDay {
		log.Fatalf("Failed to convert times all day: [%v]\n", times.Always)
	}
	if times.Text != i.Availability.Time {
		log.Fatalf("Failed to convert times text: [%v]\n", times.Text)
	}

	if len(times.Array) != len(i.Availability.TimeArray) {
		log.Fatalf("Failed to convert times array: [%v]\n", times.Array)
	}
}

func Test_ConvertMonths(t *testing.T) {
	i := sampleSimpleSourceItem()
	months := i.convertMonths()
	// only test top level, lower test will check the months per hemisphere
	if months.Always != i.Availability.IsAllYear {
		log.Fatalf("Failed to convert months always flag: [%v]\n", months.Always)
	}
}

func Test_ConvertMonthNorth(t *testing.T) {
	i := sampleSimpleSourceItem()
	months := i.convertMonthsNorth()

	if len(months.Array) != len(i.Availability.MonthArrayNorthern) {
		log.Fatalf("Failed to convert months north array: [%v]\n", months.Array)
	}

	if months.Ranges != i.Availability.MonthNorthern {
		log.Fatalf("Failed to convert months north range: [%v]\n", months.Ranges)
	}

	if months.Text != "January - February, November - December" {
		log.Fatalf("Failed to convert months north text: [%v]\n", months.Text)
	}
}

func Test_GenerateHas(t *testing.T) {
	i := sampleSimpleSourceItem()
	has := i.generateHas()

	if has.Price != true {
		log.Fatalf("Failed to generate has price: [%v]\n", has.Price)
	}
	if has.Shadow != true {
		log.Fatalf("Failed to generate has Shadow: [%v]\n", has.Shadow)
	}
	if has.Speed != false {
		log.Fatalf("Failed to generate has speed: [%v]\n", has.Speed)
	}
	if has.Rarity != true {
		log.Fatalf("Failed to generate has rarity: [%v]\n", has.Rarity)
	}
	if has.Location != true {
		log.Fatalf("Failed to generate has location: [%v]\n", has.Location)
	}
	if has.Availability != true {
		log.Fatalf("Failed to generate has availability: [%v]\n", has.Availability)
	}

}
