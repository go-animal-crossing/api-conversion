package config

// ANCHType maps types
type ANCHType int

// ANCHHemisphere maps hemispheres
type ANCHHemisphere int

// ANCHIs track new / leaving / available
type ANCHIs int

// Map for model types
const (
	NoType ANCHType = iota
	Bug
	Fish
	Sea
)

// hemisphere
const (
	NoHemisphere ANCHHemisphere = iota
	North
	South
)

// available
const (
	NoIs ANCHIs = iota
	New
	Leaving
	Available
)

// NorthHemisphere northern data
var NorthHemisphere = HemisphereItem{ID: North, Name: "Northern", Slug: "northern"}

// SouthHemisphere south data
var SouthHemisphere = HemisphereItem{ID: South, Name: "Southern", Slug: "southern"}

// Has provides meta info on what properties this item has
type Has struct {
	Price        bool `json:"price"`
	Shadow       bool `json:"shadow"`
	Speed        bool `json:"speed"`
	Rarity       bool `json:"rarity"`
	Location     bool `json:"location"`
	Availability bool `json:"availability"`
}

// ModelConfig contains model data (fish / bug etc)
type ModelConfig struct {
	Name  string   `json:"name"`
	Slug  string   `json:"slug"`
	Title string   `json:"title"`
	IsA   ANCHType `json:"is"`
	Has   Has      `json:"has"`
}

// IsItem shares `is` info for new / available
type IsItem struct {
	ID   ANCHIs `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// HemisphereItem shares info for north / south
type HemisphereItem struct {
	ID   ANCHHemisphere `json:"id"`
	Name string         `json:"name"`
	Slug string         `json:"slug"`
}

// Data provides struct for config data
type Data struct {
	APIEndpoint       string
	ModelConfigs      map[string]ModelConfig
	IsOptions         []IsItem
	HemisphereOptions []HemisphereItem
}
