package config

// Config is main configuration data source
var Config Data

func init() {

	models := map[string]ModelConfig{
		"fish": ModelConfig{
			Name:  "fish",
			Slug:  "fish",
			Title: "Fish",
			IsA:   Fish,
			Has:   Has{Price: true, Shadow: true, Speed: false, Rarity: true, Location: true, Availability: true}},
		"bugs": ModelConfig{
			Name:  "bugs",
			Slug:  "bugs",
			Title: "Bugs",
			IsA:   Bug,
			Has:   Has{Price: true, Shadow: false, Speed: false, Rarity: true, Location: true, Availability: true}},
		"sea": ModelConfig{
			Name:  "sea",
			Slug:  "sea-creatures",
			Title: "Sea Creatures",
			IsA:   Sea,
			Has:   Has{Price: true, Shadow: true, Speed: true, Rarity: false, Location: true, Availability: true}}}

	Config = Data{
		APIEndpoint:  "https://acnhapi.com/v1a/",
		ModelConfigs: models,
		IsOptions: []IsItem{
			IsItem{ID: New, Name: "New", Slug: "new"},
			IsItem{ID: Leaving, Name: "Leaving", Slug: "leaving"},
			IsItem{ID: Available, Name: "Available", Slug: "available"}},
		HemisphereOptions: []HemisphereItem{
			NorthHemisphere,
			SouthHemisphere}}
}

// API returns the api endpoint
func (c *Data) API() string {
	return c.APIEndpoint
}

// Models retuns list of endpoints
func (c *Data) Models() []string {
	models := []string{}
	for k := range c.ModelConfigs {
		models = append(models, k)
	}
	return models
}
