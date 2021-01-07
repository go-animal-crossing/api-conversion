package target

import (
	"fmt"
	"strconv"
)

// GetByID returns a slice of item matching the Ids passed
//
// - Allows multiple ids, but this does muddy waters when items arent found
// -- passing ids of 1 & 999999, if 1 is found, but not 999999, no error
//	would be returned
//
// - If cant find an id that matches, then return and error
//
// - As Id can exist in different types (fish and bugs both have a 1)
// this returns a slice of items
func (t *Target) GetByID(ids ...int) (found []Item, err error) {

	// loop over all requested ids
	for _, intID := range ids {
		id := strconv.Itoa(intID)
		// loop over all items in .All
		for _, item := range t.All {
			// compare and append if matched
			if item.ID == id {
				found = append(found, item)
			}
		}
	}
	if len(found) < 1 {
		err = fmt.Errorf("Failed to find item with ids: [%v]", ids)
	}
	return
}

// Get returns a slice of matching items based on the slug passed
//
// - Allows multiple slugs (test, bee), but if either are found then
//	no error is returned as the error is generated on a total count
//
// - slice returned maybe multiple items if matches exist between
// types (fish & bugs both having matching item)
func (t *Target) Get(slugs ...string) (found []Item, err error) {

	for _, slug := range slugs {
		for _, item := range t.All {
			if item.Slug == slug {
				found = append(found, item)
			}
		}
	}
	if len(found) < 1 {
		err = fmt.Errorf("Failed to find item with slugs [%v]", slugs)
	}
	return
}
