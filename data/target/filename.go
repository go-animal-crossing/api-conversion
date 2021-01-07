package target

import "fmt"

// Filename generates a consistent filename
func (i *Item) Filename(prefix string) string {

	return fmt.Sprintf("%s/%s/%s.json",
		prefix,
		i.Config.Slug,
		i.Slug)

}
