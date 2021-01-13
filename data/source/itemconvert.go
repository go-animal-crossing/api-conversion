package source

import (
	"acnh/data/config"
	"acnh/data/target"
	"strconv"
	"strings"
	"time"
)

// Overwrites adjusts the item data based on
// errors in the source api
func (i *Item) Overwrites() {
	// abalone has an issue with its southern data
	if i.FileName == "abalone" {
		i.Availability.MonthArraySouthern = []int{12, 1, 2, 3, 4, 5, 6, 7}
	}

}

// Convert this item to a target item
func (i *Item) Convert() (target.Item, error) {
	i.Overwrites()
	// create the target struct
	ti := target.Item{
		ID:           i.convertID(),
		Title:        i.convertName(),
		TitleSafe:    safe(i.convertName()),
		Template:     i.defaultTemplate(),
		Type:         i.convertType(),
		Prices:       i.convertPrices(),
		Images:       i.convertImages(),
		Phrases:      i.convertPhrases(),
		Attributes:   i.convertAttributes(),
		Availability: i.convertAvailability(),
		Has:          i.generateHas(),
		Config:       i.Config}

	ti.Slugify()
	ti.UpdateIs(time.Now())
	// set flag so this has been converted
	ti.Converted = true
	return ti, nil
}

// AddConfig assigns
func (i *Item) AddConfig(cnf config.ModelConfig) {
	i.Config = cnf
}

// defaultTemplate is a single page item
func (i *Item) defaultTemplate() string {
	return "page.html"
}

// ConvertID changes the default int ID to a string
func (i *Item) convertID() string {
	return strconv.Itoa(i.ID)
}

// convertName returns the EU-EN version of the name property
func (i *Item) convertName() string {
	return strings.Title(i.Names.EuEn)
}

// convertType provides target type data from this source item
func (i *Item) convertType() target.Type {
	return target.Type{
		Title: i.Config.Title,
		Slug:  i.Config.Slug,
		IsA:   i.Config.IsA}
}

// convertPrices provides target price data from this source item
func (i *Item) convertPrices() target.Prices {
	return target.Prices{
		Store: i.Price,
		Cj:    i.PriceCj,
		Flick: i.PriceFlick}
}

// convertImages provides target imagery
func (i *Item) convertImages() target.Images {
	return target.Images{
		ThumbSource: i.IconURI,
		MainSource:  i.ImageURI}
}

// convertPhrases returns target formatted struct
func (i *Item) convertPhrases() target.Phrases {
	return target.Phrases{
		Museum:      i.MuseumPhrase,
		Capture:     i.CatchPhrase,
		MuseumSafe:  safe(i.MuseumPhrase),
		CaptureSafe: safe(i.CatchPhrase)}
}

func safe(str string) string {
	str = strings.ReplaceAll(str, "'", "")
	str = strings.ReplaceAll(str, "\"", "")
	return str
}

// convertAttributes merges shadow, speed etc together
func (i *Item) convertAttributes() target.Attributes {
	return target.Attributes{
		Speed:  i.Speed,
		Shadow: i.Shadow}
}

// convertAvailability creates a target data structure with
// all Availability data stored
func (i *Item) convertAvailability() target.Availability {
	return target.Availability{
		Location: i.Availability.Location,
		Rarity:   i.Availability.Rarity,
		Months:   i.convertMonths(),
		Times:    i.convertTimes()}
}

// convertMonths creates that month data for each hemisphere
// as well as handling the always flag
func (i *Item) convertMonths() target.Months {
	return target.Months{
		Always:   i.Availability.IsAllYear,
		Northern: i.convertMonthsNorth(),
		Southern: i.convertMonthsSouth()}
}

// convertMonthsNorth handles the data swap and then via its
// extra call to MonthHemisphere struct returns the neatly
// formatted string for months (jan - feb, jul - aug etc)
func (i *Item) convertMonthsNorth() target.MonthHemisphere {
	t := target.MonthHemisphere{
		Always: i.Availability.IsAllYear,
		Ranges: i.Availability.MonthNorthern,
		Array:  i.Availability.MonthArrayNorthern,
		IsA:    config.North}
	return t.WithText()
}

// convertMonthsSouth is the same as convertMonthsNorth
func (i *Item) convertMonthsSouth() target.MonthHemisphere {
	t := target.MonthHemisphere{
		Always: i.Availability.IsAllYear,
		Ranges: i.Availability.MonthSouthern,
		Array:  i.Availability.MonthArraySouthern,
		IsA:    config.South}
	return t.WithText()
}

// convertTimes generates a target time based on source data
func (i *Item) convertTimes() target.Times {
	return target.Times{
		Always: i.Availability.IsAllDay,
		Text:   i.Availability.Time,
		Array:  i.Availability.TimeArray}
}

// generateHas simply looks at the items data and sets bool
// flags on a has struct
func (i *Item) generateHas() config.Has {

	available := ((len(i.Availability.MonthArrayNorthern) > 0) ||
		(len(i.Availability.MonthArraySouthern) > 0))

	return config.Has{
		Price:        (i.Price > 0),
		Shadow:       (len(i.Shadow) > 0),
		Speed:        (len(i.Speed) > 0),
		Rarity:       (len(i.Availability.Rarity) > 0),
		Location:     (len(i.Availability.Location) > 0),
		Availability: available}

}
