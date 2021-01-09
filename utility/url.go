package utility

import (
	"fmt"
	"strings"
)

// Slugify converts string passed to a url friendly slug
func Slugify(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "-")
	str = strings.ReplaceAll(str, "'", "")
	return str
}

// URL converts a series of strings to a / separated string of slugs
func URL(strs ...string) string {

	url := ""
	for _, str := range strs {
		url = fmt.Sprintf("%s/%s", url, Slugify(str))
	}
	return url
}
