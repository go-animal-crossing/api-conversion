package pages

import (
	"acnh/data/config"
	"acnh/data/target"
)

// Page is main data container
//
// IsList
// A flag to work out if this is a list of many items (such as homepage)
// or a item page (like pale chub)
//
// Items
// This would be pulled from Target.All for listing pages
//
type Page struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	H1       string        `json:"h1"`
	Template string        `json:"template"`
	URI      string        `json:"url"`
	Filename string        `json:"-"` //exclude
	Type     target.Type   `json:"type"`
	IsList   bool          `json:"list"`
	Items    []target.Item `json:"items,omitempty"`
	Has      config.Has    `json:"has"`
	Meta     Meta          `json:"meta,omitempty"`
	Grid     Grid          `json:"grid,omitempty"`
}

// Grid is used for home page and sharing pages
type Grid struct {
	Type       bool `json:"type"`
	Is         bool `json:"is"`
	Hemisphere bool `json:"hemisphere"`
}

// Meta contains top level info for listing
type Meta struct {
	Type       string `json:"type"`
	Is         string `json:"is"`
	Hemisphere string `json:"hemisphere"`
	Month      string `json:"month"`
	ItemCount  int    `json:"counter"`
	Links      Links  `json:"links"`
}

// Links provides info for related links on list pages
// - view /north/x etc
type Links struct {
	Related string      `json:"related"`
	Months  []MonthLink `json:"months"`
}

// MonthLink provides data info
type MonthLink struct {
	MonthShort string `json:"month_short"`
	MonthLong  string `json:"month_long"`
}
