package target

import (
	"log"
	"testing"
)

func Test_GetByID(t *testing.T) {

	tg := Target{}

	// items should have two items
	items := sampleMultiTargetItems()
	if len(items) != 2 {
		log.Fatalf("Item setup failure!\n%v\n", items)
	}
	// assign items
	tg.All = items

	// this id is not in the data set, so should fail and found should be empty
	found, e := tg.GetByID(-1)
	if e == nil || len(found) > 0 {
		log.Fatalf("Found an item by id (-1) when nothing should be set:\n %v\n%v\n", found, e)
	}

	// should only find 1
	found, e = tg.GetByID(1)
	if len(found) != 1 || e != nil {
		log.Fatalf("Failed to find an item by id (1) when 1 item should be:\n %v\n%v\n", found, e)
	}
	found, e = tg.GetByID(2)
	if len(found) != 1 || e != nil {
		log.Fatalf("Failed to find an item by id (2) when 1 item should be:\n %v\n%v\n", found, e)
	}

	// test getting more than 1 id, but only one exists
	found, e = tg.GetByID(1, -1)
	if len(found) != 1 || e != nil {
		log.Fatalf("Failed to find an item by multiple ids (1, -1) when 1 item should be:\n %v\n%v\n", found, e)
	}

	// find 2 items correctly
	found, e = tg.GetByID(1, 2)
	if len(found) != 2 || e != nil {
		log.Fatalf("Failed to find items by multiple ids (1, 2) when 2 item should be:\n %v\n%v\n", found, e)
	}

	// now we add items again to T to ensure we find 2 items with the same id
	tg.All = append(tg.All, items...)
	found, e = tg.GetByID(1)
	if len(found) != 2 || e != nil {
		log.Fatalf("Failed to find items by id (1) when 2 item should be:\n %v\n%v\n", found, e)
	}
}

func Test_Get(t *testing.T) {

	tg := Target{}

	// items should have two items
	items := sampleMultiTargetItems()
	if len(items) != 2 {
		log.Fatalf("Item setup failure!\n%v\n", items)
	}
	// assign items!
	tg.All = items

	// this is not in the data set, so should fail and found should be empty
	found, e := tg.Get("FOOBAR")
	if e == nil || len(found) > 0 {
		log.Fatalf("Found an item by slug (FOOBAR) when nothing should be set:\n %v\n%v\n%v\n", len(found), found, e)
	}
	// get a single real version
	found, e = tg.Get("pale-chub")
	if e != nil || len(found) != 1 {
		log.Fatalf("Failed to find an item by slug (pale-chub) when 1 item should be:\n %v\n%v\n%v\n", len(found), found, e)
	}
	found, e = tg.Get("pale-chub", "foobar")
	if e != nil || len(found) != 1 {
		log.Fatalf("Failed to find an item by slug (pale-chub, foobar) when 1 item should be:\n %v\n%v\n%v\n", len(found), found, e)
	}

	// get two real versions
	found, e = tg.Get("pale-chub", "bitterling")
	if e != nil || len(found) != 2 {
		log.Fatalf("Failed to find items by slug (pale-chub, bitterling) when 2 item should be:\n %v\n%v\n%v\n", len(found), found, e)
	}

}
