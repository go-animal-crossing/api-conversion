package target

import (
	"acnh/data/config"
	"encoding/json"
)

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

// sampleNewInFebTest contains 3 items, 2 that are new for Feb, 1 that isnt
func sampleNewInFebTest() []Item {

	raw := `[
		{
			"id": "12",
			"title": "Soft-Shelled Turtle",
			"title_safe": "Soft-Shelled Turtle",
			"url": "/fish/soft-shelled-turtle",
			"slug": "soft-shelled-turtle",
			"template": "page.html",
			"type": {
				"title": "Fish",
				"slug": "fish",
				"is": 2
			},
			"prices": {
				"store": 3750,
				"cj": 5625,
				"flick": 0
			},
			"phrases": {
				"capture": "I caught a soft-shelled turtle! I can really shell it out!",
				"capture_safe": "I caught a soft-shelled turtle! I can really shell it out!",
				"museum": "These relatives of common turtles will not let go when they decide to bite down on something. Although they can be quite shy, they will become a bit aggressive when threatened. They are often found in water and are very good swimmers because of the flat shape of their shells. They swim with their snouts over the surface of the water so they can breathe. Soft-shelled turtles aren't safe to hold because they are so prone to biting at the slightest movement. However, if one bites down on you, it will usually let go when you put it safely back in the water.",
				"museum_safe": "These relatives of common turtles will not let go when they decide to bite down on something. Although they can be quite shy, they will become a bit aggressive when threatened. They are often found in water and are very good swimmers because of the flat shape of their shells. They swim with their snouts over the surface of the water so they can breathe. Soft-shelled turtles arent safe to hold because they are so prone to biting at the slightest movement. However, if one bites down on you, it will usually let go when you put it safely back in the water."
			},
			"images": {
				"thumb_source": "https://acnhapi.com/v1/icons/fish/12",
				"thumb": "/images/soft-shelled-turtle/thumb.png",
				"main_source": "https://acnhapi.com/v1/images/fish/12",
				"main": "/images/soft-shelled-turtle/main.png"
			},
			"availability": {
				"location": "River",
				"rarity": "Uncommon",
				"months": {
					"always": false,
					"northern": {
						"always": false,
						"text": "August - September",
						"ranges": "8-9",
						"array": [
							8,
							9
						],
						"sequences": [
							[
								8,
								9
							]
						],
						"is": 1
					},
					"southern": {
						"always": false,
						"text": "February - March",
						"ranges": "2-3",
						"array": [
							2,
							3
						],
						"sequences": [
							[
								2,
								3
							]
						],
						"is": 2
					}
				},
				"times": {
					"always": false,
					"text": "4pm - 9am",
					"array": [
						16,
						17,
						18,
						19,
						20,
						21,
						22,
						23,
						0,
						1,
						2,
						3,
						4,
						5,
						6,
						7,
						8
					]
				}
			},
			"attributes": {
				"shadow": "Medium (4)",
				"speed": ""
			},
			"has": {
				"price": true,
				"shadow": true,
				"speed": false,
				"rarity": true,
				"location": true,
				"availability": true
			},
			"is": {
				"time_basedon": "2021-02-01T00:00:00Z",
				"northern": {
					"time_basedon": "2021-02-01T00:00:00Z",
					"new": false,
					"leaving": false,
					"available": false
				},
				"southern": {
					"time_basedon": "2021-02-01T00:00:00Z",
					"new": true,
					"leaving": false,
					"available": true
				}
			},
			"typeConfig": {
				"name": "fish",
				"slug": "fish",
				"title": "Fish",
				"is": 2,
				"has": {
					"price": true,
					"shadow": true,
					"speed": false,
					"rarity": true,
					"location": true,
					"availability": true
				}
			},
			"converted": true
		},
		{
			"id": "44",
			"title": "Tiger Beetle",
			"title_safe": "Tiger Beetle",
			"url": "/bugs/tiger-beetle",
			"slug": "tiger-beetle",
			"template": "page.html",
			"type": {
				"title": "Bugs",
				"slug": "bugs",
				"is": 1
			},
			"prices": {
				"store": 1500,
				"cj": 0,
				"flick": 2250
			},
			"phrases": {
				"capture": "I caught a tiger beetle! I pounced first!",
				"capture_safe": "I caught a tiger beetle! I pounced first!",
				"museum": "The tiger beetle is extremely fleet of foot, though it runs in a rather peculiar way. That is, it sprints, then stops, then sprints again...using these speedy maneuvers to run down its prey. You see, the tiger beetle—like a real tiger—is a powerful predator. The mere thought of it giving chase gives me the willies. Tigers are terrifying at any size.",
				"museum_safe": "The tiger beetle is extremely fleet of foot, though it runs in a rather peculiar way. That is, it sprints, then stops, then sprints again...using these speedy maneuvers to run down its prey. You see, the tiger beetle—like a real tiger—is a powerful predator. The mere thought of it giving chase gives me the willies. Tigers are terrifying at any size."
			},
			"images": {
				"thumb_source": "https://acnhapi.com/v1/icons/bugs/44",
				"thumb": "/images/tiger-beetle/thumb.png",
				"main_source": "https://acnhapi.com/v1/images/bugs/44",
				"main": "/images/tiger-beetle/main.png"
			},
			"availability": {
				"location": "On the ground",
				"rarity": "Uncommon",
				"months": {
					"always": false,
					"northern": {
						"always": false,
						"text": "February - November",
						"ranges": "2-11",
						"array": [
							2,
							3,
							4,
							5,
							6,
							7,
							8,
							9,
							10,
							11
						],
						"sequences": [
							[
								2,
								3,
								4,
								5,
								6,
								7,
								8,
								9,
								10,
								11
							]
						],
						"is": 1
					},
					"southern": {
						"always": false,
						"text": "August - April",
						"ranges": "8-4",
						"array": [
							8,
							9,
							10,
							11,
							12,
							1,
							2,
							3,
							4
						],
						"sequences": [
							[
								8,
								9,
								10,
								11,
								12,
								1,
								2,
								3,
								4
							]
						],
						"is": 2
					}
				},
				"times": {
					"always": true,
					"text": "",
					"array": [
						0,
						1,
						2,
						3,
						4,
						5,
						6,
						7,
						8,
						9,
						10,
						11,
						12,
						13,
						14,
						15,
						16,
						17,
						18,
						19,
						20,
						21,
						22,
						23
					]
				}
			},
			"attributes": {
				"shadow": "",
				"speed": ""
			},
			"has": {
				"price": true,
				"shadow": false,
				"speed": false,
				"rarity": true,
				"location": true,
				"availability": true
			},
			"is": {
				"time_basedon": "2021-02-01T00:00:00Z",
				"northern": {
					"time_basedon": "2021-02-01T00:00:00Z",
					"new": true,
					"leaving": false,
					"available": true
				},
				"southern": {
					"time_basedon": "2021-02-01T00:00:00Z",
					"new": false,
					"leaving": false,
					"available": true
				}
			},
			"typeConfig": {
				"name": "bugs",
				"slug": "bugs",
				"title": "Bugs",
				"is": 1,
				"has": {
					"price": true,
					"shadow": false,
					"speed": false,
					"rarity": true,
					"location": true,
					"availability": true
				}
			},
			"converted": true
		},
		{
            "id": "58",
            "title": "Barred Knifejaw",
            "title_safe": "Barred Knifejaw",
            "url": "/fish/barred-knifejaw",
            "slug": "barred-knifejaw",
            "template": "page.html",
            "type": {
                "title": "Fish",
                "slug": "fish",
                "is": 2
            },
            "prices": {
                "store": 5000,
                "cj": 7500,
                "flick": 0
            },
            "phrases": {
                "capture": "I caught a barred knifejaw! They must have a hard time eating!",
                "capture_safe": "I caught a barred knifejaw! They must have a hard time eating!",
                "museum": "Wonderful! I would love nothing more! The barred knifejaw, also called the striped beakfish, is a lovely fish with a distinctive striped pattern. It is reportedly also a very curious fish in that it has been known to approach swimmers to \"greet\" them. They also tend to be good at avoiding fishing hooks, so well done on outwitting this one!",
                "museum_safe": "Wonderful! I would love nothing more! The barred knifejaw, also called the striped beakfish, is a lovely fish with a distinctive striped pattern. It is reportedly also a very curious fish in that it has been known to approach swimmers to greet them. They also tend to be good at avoiding fishing hooks, so well done on outwitting this one!"
            },
            "images": {
                "thumb_source": "https://acnhapi.com/v1/icons/fish/58",
                "thumb": "/images/barred-knifejaw/thumb.png",
                "main_source": "https://acnhapi.com/v1/images/fish/58",
                "main": "/images/barred-knifejaw/main.png"
            },
            "availability": {
                "location": "Sea",
                "rarity": "Uncommon",
                "months": {
                    "always": false,
                    "northern": {
                        "always": false,
                        "text": "March - November",
                        "ranges": "3-11",
                        "array": [
                            3,
                            4,
                            5,
                            6,
                            7,
                            8,
                            9,
                            10,
                            11
                        ],
                        "sequences": [
                            [
                                3,
                                4,
                                5,
                                6,
                                7,
                                8,
                                9,
                                10,
                                11
                            ]
                        ],
                        "is": 1
                    },
                    "southern": {
                        "always": false,
                        "text": "September - May",
                        "ranges": "9-5",
                        "array": [
                            9,
                            10,
                            11,
                            12,
                            1,
                            2,
                            3,
                            4,
                            5
                        ],
                        "sequences": [
                            [
                                9,
                                10,
                                11,
                                12,
                                1,
                                2,
                                3,
                                4,
                                5
                            ]
                        ],
                        "is": 2
                    }
                },
                "times": {
                    "always": true,
                    "text": "",
                    "array": [
                        0,
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7,
                        8,
                        9,
                        10,
                        11,
                        12,
                        13,
                        14,
                        15,
                        16,
                        17,
                        18,
                        19,
                        20,
                        21,
                        22,
                        23
                    ]
                }
            },
            "attributes": {
                "shadow": "Medium (3)",
                "speed": ""
            },
            "has": {
                "price": true,
                "shadow": true,
                "speed": false,
                "rarity": true,
                "location": true,
                "availability": true
            },
            "is": {
                "time_basedon": "2021-09-01T00:00:00Z",
                "northern": {
                    "time_basedon": "2021-09-01T00:00:00Z",
                    "new": false,
                    "leaving": false,
                    "available": true
                },
                "southern": {
                    "time_basedon": "2021-09-01T00:00:00Z",
                    "new": true,
                    "leaving": false,
                    "available": true
                }
            },
            "typeConfig": {
                "name": "fish",
                "slug": "fish",
                "title": "Fish",
                "is": 2,
                "has": {
                    "price": true,
                    "shadow": true,
                    "speed": false,
                    "rarity": true,
                    "location": true,
                    "availability": true
                }
            },
            "converted": true
        }
	]`

	multi := make([]Item, 0)
	j := []byte(raw)
	json.Unmarshal(j, &multi)

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
