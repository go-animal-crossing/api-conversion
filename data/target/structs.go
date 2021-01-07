package target

import (
	"acnh/data/config"
	"time"
)

// Filter struct
// - Month has no effect unless coupled with another
//   field - it is simply used to trigger regen of Is data
type Filter struct {
	Is         config.ANCHIs
	Hemisphere config.ANCHHemisphere
	Month      time.Month
	Type       config.ANCHType
}

// Target provides
type Target struct {
	All       []Item
	Converted bool
}

// Item is a target structure
type Item struct {
	ID           string             `json:"id"`
	Title        string             `json:"title"`
	URI          string             `json:"url"`
	Slug         string             `json:"slug"`
	Template     string             `json:"template"`
	Type         Type               `json:"type"`
	Prices       Prices             `json:"prices"`
	Phrases      Phrases            `json:"phrases"`
	Images       Images             `json:"images"`
	Availability Availability       `json:"availability"`
	Attributes   Attributes         `json:"attributes"`
	Has          config.Has         `json:"has"`
	Is           Is                 `json:"is"`
	Config       config.ModelConfig `json:"typeConfig"`
	Converted    bool               `json:"converted"`
}

// Is provides meta data on if this item is new, leaving, etc
type Is struct {
	Time     time.Time `json:"time_basedon"`
	Northern IsConfig  `json:"northern"`
	Southern IsConfig  `json:"southern"`
}

// IsConfig provides boolean data for a hempishere to determine if this
// item is new for that that hemisphere this period
type IsConfig struct {
	Time       time.Time `json:"time_basedon"`
	New        bool      `json:"new"`
	Leaving    bool      `json:"leaving"`
	Availabile bool      `json:"available"`
}

// Attributes provides some meta fields (speed, shadow etc)
type Attributes struct {
	Shadow string `json:"shadow"`
	Speed  string `json:"speed"`
}

// Availability provides
type Availability struct {
	Location string `json:"location"`
	Rarity   string `json:"rarity"`
	Months   Months `json:"months"`
	Times    Times  `json:"times"`
}

// Months provides structured data for Availability
type Months struct {
	Always   bool            `json:"always"`
	Northern MonthHemisphere `json:"northern"`
	Southern MonthHemisphere `json:"southern"`
}

// MonthHemisphere provides hemisphere specific data
type MonthHemisphere struct {
	Always    bool                  `json:"always"`
	Text      string                `json:"text"`
	Ranges    string                `json:"ranges"`
	Array     []int                 `json:"array"`
	Sequences [][]int               `json:"sequences"`
	IsA       config.ANCHHemisphere `json:"is"`
}

// Times provides times based Availability data
type Times struct {
	Always bool   `json:"always"`
	Text   string `json:"text"`
	Array  []int  `json:"array"`
}

// Type provides
type Type struct {
	Title string          `json:"title"`
	Slug  string          `json:"slug"`
	IsA   config.ANCHType `json:"is"`
}

// Prices provides
type Prices struct {
	Store int `json:"store"`
	Cj    int `json:"cj"`
	Flick int `json:"flick"`
}

// Images provide
type Images struct {
	Thumb string `json:"thumb"`
	Main  string `json:"main"`
}

// Phrases provides
type Phrases struct {
	Capture string `json:"capture"`
	Museum  string `json:"museum"`
}
