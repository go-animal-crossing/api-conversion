package source

import "acnh/data/config"

func sampleSimpleSourceItem() Item {

	return Item{
		ID:           1,
		Shadow:       "Smallest (1)",
		Price:        900,
		PriceFlick:   1350,
		CatchPhrase:  "Yes!",
		MuseumPhrase: "Found it!",
		ImageURI:     "http://g.com",
		IconURI:      "https://g.co.uk",
		Names: Name{
			EuEn: "bitterling",
			UsEn: "Test"},
		Availability: Availability{
			MonthNorthern:      "1-2,11-12",
			MonthSouthern:      "6-7",
			MonthArrayNorthern: []int{1, 2, 11, 12},
			MonthArraySouthern: []int{6, 7},
			Time:               "9-5",
			TimeArray:          []int{9, 10, 11, 12, 13, 14, 15, 16, 17},
			IsAllDay:           true,
			IsAllYear:          false,
			Location:           "River",
			Rarity:             "Common"},
		// use fish config
		Config: config.Config.ModelConfigs["fish"]}

}
