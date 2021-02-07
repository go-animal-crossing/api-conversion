package target

import (
	"acnh/data/config"
	"acnh/utility"
)

// Filter uses the configured Filter struct to find
//
// - Filter.Month by itself wont have an impact
// it triggers an update to Is data struct to allow
// parsing for available etc
//
// Item.Is struct data is generated using time
// based context - which is why its rengenerated
// if we change the Month being looked at via
// via the filter
//
// - If an empty filter is passed, returns all
func (t *Target) Filter(filter Filter) (found []Item) {

	found = t.All

	// if month has been filtered, we now need
	// to update all Is data
	if mnth := int(filter.Month); mnth > 0 {
		ts := utility.MonthToTime(mnth)
		t.UpdateIs(&found, ts)
	}

	// Filter by type outside of other items as thats top level
	// restriction
	if ty := filter.Type; ty > 0 {
		found = t.filterByType(ty, found)
	}

	// if hemisphere is set, then we need to restrict data to
	// what matches
	if h := filter.Hemisphere; h > 0 {
		// if Is filter is set, then we filter by both hemisphere
		// and Is
		// otherwise, just check hemisphere as before, which is
		// effectively Available
		if is := filter.Is; is > 0 {
			found = t.filterByIsAndHemisphere(is, h, found)
		} else {
			found = t.filterByHemisphere(h, found)
		}

	} else {
		// if no hemisphere is set, then just check on the
		// the is, otherwise dont filter (to preserve empty filter returning all)
		if is := filter.Is; is > 0 {
			found = t.filterByIs(is, found)
		}

	}

	return
}

// filterByHemisphere returns only items that available for h set to true
// (as new / leaving would also require available to be true)
func (t *Target) filterByHemisphere(h config.ANCHHemisphere, all []Item) (found []Item) {
	found = make([]Item, 0)
	for _, item := range all {
		// work out which IS config to check
		var check IsConfig
		switch h {
		case config.South:
			check = item.Is.Southern
		default:
			check = item.Is.Northern
		}

		if check.Availabile {
			found = append(found, item)
		}
	}
	return
}

// filterByIsAndHemisphere matches on both Is and Hemishphere
func (t *Target) filterByIsAndHemisphere(is config.ANCHIs, h config.ANCHHemisphere, all []Item) (found []Item) {
	found = make([]Item, 0)
	for _, item := range all {
		// determine which hemisphere to use
		hemi := item.Is.Northern
		if h == config.SouthHemisphere.ID {
			hemi = item.Is.Southern
		}
		// get all is data for this item in map
		mapped := map[config.ANCHIs]bool{
			config.Leaving:   hemi.Leaving,
			config.Available: hemi.Availabile,
			config.New:       hemi.New}
		// compare the index, if that is true, then add
		if mapped[is] == true {
			found = append(found, item)
		}
	}

	return
}

// filterByIs reduces all to only those that match new / available / leaving
func (t *Target) filterByIs(is config.ANCHIs, all []Item) (found []Item) {
	found = make([]Item, 0)
	for _, item := range all {
		// get all is data for this item in map
		mapped := map[config.ANCHIs]bool{
			config.Leaving:   (item.Is.Northern.Leaving || item.Is.Southern.Leaving),
			config.Available: (item.Is.Northern.Availabile || item.Is.Southern.Availabile),
			config.New:       (item.Is.Northern.New || item.Is.Southern.New)}
		// compare the index, if that is true, then add
		if mapped[is] == true {
			found = append(found, item)
		}
	}

	return
}

// filterByType reduces all to only those items with Type.IsA matching value
func (t *Target) filterByType(ty config.ANCHType, all []Item) (found []Item) {

	// filter to just this type of item
	found = make([]Item, 0)
	for _, item := range all {
		if item.Type.IsA == ty {
			found = append(found, item)
		}
	}
	return

}
