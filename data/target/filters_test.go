package target

import (
	"acnh/data/config"
	"log"
	"testing"
	"time"
)

func Test_Filters_SimpleType(t *testing.T) {

	tg := Target{}
	// items should have two items
	items := sampleMultiTargetItems()
	if len(items) != 2 {
		log.Fatalf("Item setup failure!\n%v\n", items)
	}
	// assign items to fish!
	tg.All = items
	// filter call with empty filter
	f := tg.Filter(Filter{})
	if len(f) != 2 {
		log.Fatalf("Empty filter call failed, should return 2, returned [%v]\n%v\n", len(f), f)
	}

	// filter call for fish
	f = tg.Filter(Filter{Type: config.Fish})
	if len(f) != 2 {
		log.Fatalf("Fish filter call failed, should return 2, returned [%v]\n%v\n", len(f), f)
	}
	// filter call with sea creatures
	f = tg.Filter(Filter{Type: config.Sea})
	if len(f) != 0 {
		log.Fatalf("Filter by sea call failed, should return 0, returned [%v]\n%v\n", len(f), f)
	}

	// filter call with month and nothing else should have no effect
	f = tg.Filter(Filter{Month: time.Month(1)})
	if len(f) != 2 {
		log.Fatalf("Just month filter call failed, should return 2, returned [%v]\n%v\n", len(f), f)
	}

}

func Test_Filters_TypeIsCombinations(t *testing.T) {
	tg := Target{}
	// items should have two items
	items := sampleMultipleVariedTargetItems()
	if len(items) != 10 {
		log.Fatalf("Item setup failure!\n%v\n", items)
	}
	tg.All = items

	// FISH TESTS
	// fish = 4
	f := tg.Filter(Filter{Type: config.Fish})
	if len(f) != 4 {
		log.Fatalf("Fish filter call failed, should return 4, returned [%v]\n%v\n", len(f), f)
	}
	// fish that are new
	f = tg.Filter(Filter{Type: config.Fish, Is: config.New})
	if len(f) != 1 {
		log.Fatalf("Fish & New filter call failed, should return 1, returned [%v]\n%v\n", len(f), f)
	}
	// fish that are leaving
	f = tg.Filter(Filter{Type: config.Fish, Is: config.Leaving})
	if len(f) != 2 {
		log.Fatalf("Fish & Leaving filter call failed, should return 1, returned [%v]\n%v\n", len(f), f)
	}
	// fish that are leaving in the north
	f = tg.Filter(Filter{Type: config.Fish, Is: config.Leaving, Hemisphere: config.North})
	if len(f) != 2 {
		log.Fatalf("Fish & Leaving & Northen filter call failed, should return 1, returned [%v]\n%v\n", len(f), f)
	}
	// fish in the south (should be 0)
	f = tg.Filter(Filter{Type: config.Fish, Hemisphere: config.South})
	if len(f) != 0 {
		log.Fatalf("Fish & Southern filter call failed, should return 0, returned [%v]\n%v\n", len(f), f)
	}

	// BUGS
	f = tg.Filter(Filter{Type: config.Bug})
	if len(f) != 5 {
		log.Fatalf("Bug filter call failed, should return 5, returned [%v]\n%v\n", len(f), f)
	}

	// Hemisphere
	f = tg.Filter(Filter{Hemisphere: config.South})
	if len(f) != 1 {
		log.Fatalf("Hemisphere filter call failed, should return 1, returned [%v]\n%v\n", len(f), f)
	}

	// IS
	f = tg.Filter(Filter{Is: config.New})
	if len(f) != 3 {
		log.Fatalf("New filter call failed, should return 3, returned [%v]\n%v\n", len(f), f)
	}

	f = tg.Filter(Filter{Is: config.Leaving})
	if len(f) != 3 {
		log.Fatalf("Leaving filter call failed, should return 3, returned [%v]\n%v\n", len(f), f)
	}
	f = tg.Filter(Filter{Is: config.Available})
	if len(f) != 6 {
		log.Fatalf("Available filter call failed, should return 6, returned [%v]\n%v\n", len(f), f)
	}

}

func Test_Filters_Month(t *testing.T) {
	tg := Target{}
	// items should have two items
	items := sampleMultiTargetItems()
	if len(items) != 2 {
		log.Fatalf("Item setup failure!\n%v\n", items)
	}
	tg.All = items

	// august should fine only one result
	m := time.Month(8)
	f := tg.Filter(Filter{Is: config.Available, Month: m})
	if len(f) != 1 {
		log.Fatalf("Month filter call failed, should return 1, returned [%v]\n%v\n", len(f), f)
	}

	// feb, for south, should also return 1 only
	m = time.Month(2)
	f = tg.Filter(Filter{Is: config.Available, Month: m, Hemisphere: config.South})
	if len(f) != 1 {
		log.Fatalf("Month filter for south failed, should return 1, returned [%v]\n%v\n", len(f), f)
	}
}

func Test_Filters_Month_ConflictingDates(t *testing.T) {
	tg := Target{}
	// items should have two items
	items := sampleNewInFebTest()
	if len(items) != 3 {
		log.Fatalf("Item setup failure!\n%v\n", items)
	}
	tg.All = items

	m := time.Month(2)
	f := tg.Filter(Filter{Is: config.New, Month: m, Hemisphere: config.South})

	if len(f) != 1 {
		log.Fatalf("Conflicting month for South filter for south failed, should return 1, returned [%v]\n%v\n", len(f), f)
	}

	if f[0].ID != "12" {
		log.Fatalf("Filter returned incorrect item, returned:\n%v\n", f[0])
	}

	f = tg.Filter(Filter{Is: config.New, Month: m, Hemisphere: config.North})
	if len(f) != 1 {
		log.Fatalf("Conflicting month for North filter failed, should return 1, returned [%v]\n%v\n", len(f), f)
	}
}
