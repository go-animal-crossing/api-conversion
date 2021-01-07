package target

import "time"

// UpdateIs triggers the item.UpdateIs for everything in all
// passed
func (t *Target) UpdateIs(all *[]Item, ts time.Time) {

	for i := range *all {
		(*all)[i].UpdateIs(ts)
	}

}

// UpdateIs uses the date passed in to work out
// the .Is values for this item
func (i *Item) UpdateIs(ts time.Time) {
	i.Is.Time = ts
	i.Is.Northern = i.Availability.Months.Northern.Is(ts)
	i.Is.Southern = i.Availability.Months.Southern.Is(ts)

}
