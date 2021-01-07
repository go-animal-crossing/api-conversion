package source

import (
	"acnh/data/target"
)

// Convert generates a Target data set from the current Source object
func (s *Source) Convert() (target.Target, error) {

	// create the target data structures
	t := target.Target{}
	converted := make([]target.Item, 0)
	// loop over each item in the All list of source
	for _, item := range s.All {
		// convert the item over
		c, e := item.Convert()
		// if theres an error converting, return
		if e != nil {
			return t, e

		}
		// set converted items of type k to be this slice
		converted = append(converted, c)
	}
	// set the target All to the map of converted items
	t.All = converted
	// set flag to acknowledge the conversion is complete
	t.Converted = true
	return t, nil
}
