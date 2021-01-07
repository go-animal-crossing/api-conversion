package target

import "acnh/data/config"

func sampleSimpleTargetItem() Item {

	return Item{
		ID:       "1",
		Title:    "Bitterling",
		URI:      "fish/bitterling",
		Slug:     "bitterling",
		Template: "page.html",
		Type: Type{
			Title: "Fish",
			Slug:  "fish",
			IsA:   config.Fish},
		Prices: Prices{
			Store: 900,
			Cj:    1350,
			Flick: 0},
		Phrases: Phrases{
			Capture: "Caught!",
			Museum:  "Donated!"},
		Images: Images{
			Main:  "http://main.com",
			Thumb: "http://thumb.com"},
		Availability: Availability{
			Location: "River",
			Rarity:   "Common",
			Months: Months{
				Always: false,
				Northern: MonthHemisphere{
					Always:    false,
					Text:      "November - Jan",
					Ranges:    "11-1",
					Array:     []int{11, 12, 1},
					Sequences: [][]int{[]int{11, 12, 1}}},
				Southern: MonthHemisphere{
					Always:    false,
					Text:      "April - July",
					Ranges:    "4-7",
					Array:     []int{4, 5, 6, 7},
					Sequences: [][]int{[]int{4, 5, 6, 7}}}},
			Times: Times{
				Always: true,
				Text:   "",
				Array: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
					12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}}},
		Attributes: Attributes{
			Speed:  "",
			Shadow: "Smallest (1)"},
		Has: config.Has{
			Price:        true,
			Shadow:       true,
			Speed:        false,
			Rarity:       true,
			Location:     true,
			Availability: true},
		Is: Is{
			Northern: IsConfig{
				New:        false,
				Leaving:    false,
				Availabile: true},
			Southern: IsConfig{
				New:        false,
				Leaving:    false,
				Availabile: false}},
		Config: config.ModelConfig{
			Name:  "fish",
			Slug:  "fish",
			Title: "Fish",
			IsA:   config.Fish},
		Converted: true}
}

func sampleMultiTargetItems() []Item {
	one := sampleSimpleTargetItem()
	// use a second, but change id
	two := sampleSimpleTargetItem()
	two.ID = "2"
	two.Title = "Pale Chub"
	two.Slug = "pale-chub"
	// reset the Availability info to one that varies
	// against the first example for testing
	two.Availability.Months.Always = false
	two.Availability.Months.Northern = MonthHemisphere{
		Always:    false,
		Text:      "August",
		Ranges:    "8",
		Array:     []int{8},
		Sequences: [][]int{[]int{8}},
		IsA:       config.North}

	two.Availability.Months.Southern = MonthHemisphere{
		Always:    false,
		Text:      "February",
		Ranges:    "2",
		Array:     []int{2},
		Sequences: [][]int{[]int{2}},
		IsA:       config.South}

	multi := make([]Item, 0)
	multi = append(multi, one, two)

	return multi
}

// sampleMultipleVariedTargetItems data set:
// 10 items - 9 Northern, 1 Southern
// - /{fish}/ = 4
// - /{fish}/{new}/ = 1
// - /{fish}/{leaving}/ = 2
// - /{fish}/{available}/ = 2

// - /{bugs}/ = 5
// - /{bugs}/{new}/ = 2
// - /{bugs}/{leaving}/ = 0
// - /{bugs}/{available}/ = 3

// - /{sea-creatures}/ = 1 (SOUTH)
// - /{sea-creatures}/{new}/ = 0
// - /{sea-creatures}/{leaving}/ = 1
// - /{sea-creatures}/{available}/ = 1

// - /{new}/ = 3
// - /{leaving}/ = 3
// - /{available}/ = 6
func sampleMultipleVariedTargetItems() []Item {

	// FISH (already set)
	one := sampleSimpleTargetItem()
	one.Is.Northern.New = true
	one.Is.Northern.Leaving = true

	two := sampleSimpleTargetItem()
	two.ID = "2"
	two.Title = "Pale Chub"
	two.Slug = "pale-chub"
	two.Is.Northern.Leaving = true

	three := sampleSimpleTargetItem()
	three.ID = "3"
	three.Title = "three"
	three.Slug = "three"
	three.Is.Northern.Availabile = false

	four := sampleSimpleTargetItem()
	four.ID = "4"
	four.Title = "four"
	four.Slug = "four"
	four.Is.Northern.Availabile = false

	// BUGS
	five := sampleSimpleTargetItem()
	five.Type.IsA = config.Bug
	five.Is.Northern.New = true

	six := sampleSimpleTargetItem()
	six.Type.IsA = config.Bug
	six.ID = "2"
	six.Title = "Six"
	six.Slug = "six"
	six.Is.Northern.New = true

	seven := sampleSimpleTargetItem()
	seven.Type.IsA = config.Bug
	seven.ID = "3"
	seven.Title = "Seven"
	seven.Slug = "seven"

	eight := sampleSimpleTargetItem()
	eight.Type.IsA = config.Bug
	eight.ID = "4"
	eight.Title = "eight"
	eight.Slug = "eight"
	eight.Is.Northern.Availabile = false

	nine := sampleSimpleTargetItem()
	nine.Type.IsA = config.Bug
	nine.ID = "5"
	nine.Title = "nine"
	nine.Slug = "nine"
	nine.Is.Northern.Availabile = false

	// SEA
	ten := sampleSimpleTargetItem()
	ten.Type.IsA = config.Sea
	ten.Is.Northern.Availabile = false
	ten.Is.Southern.Availabile = true
	ten.Is.Southern.Leaving = true

	multi := make([]Item, 0)
	multi = append(multi, one, two, three, four, five, six, seven, eight, nine, ten)
	return multi

}
