package target

import (
	"acnh/utility"
)

// Slugify creates Slug & URI fields from the title & config
func (i *Item) Slugify() {
	i.Slug = utility.Slugify(i.Title)
	i.URI = utility.URL(i.Config.Slug, i.Title)
}
