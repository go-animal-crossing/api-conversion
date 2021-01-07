package pages

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/afero"
)

// Sort alphabtically the items slice
func (pg *Page) Sort() {

	it := pg.Items
	sort.SliceStable(it, func(i, j int) bool {
		return it[i].Title < it[j].Title
	})
	pg.Items = it
}

// SaveToFile saves contents of the page to a file
func (pg *Page) SaveToFile(fs afero.Fs, filename string) error {
	pg.Sort()
	// convert self into json string
	content, err := json.Marshal(pg)
	if err != nil {
		return err
	}
	// check and create the directory
	fileDir := filepath.Dir(filename)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		fs.MkdirAll(fileDir, os.ModePerm)
	}
	//fmt.Printf("saving to file: %v\n", filename)
	return afero.WriteFile(fs, filename, content, os.FileMode(int(0777)))
}

func (pg *Page) generateFilename(prefix string) string {
	return fmt.Sprintf("%s%s/index.json",
		prefix,
		pg.URI)
}

// Save will check .Item & .Items for length and then work our correct
// way to save that data
func (pg *Page) Save(fs afero.Fs, prefix string) error {

	if pg.IsList {
		pg.Meta.ItemCount = len(pg.Items)
		pg.Filename = pg.generateFilename(prefix)
		return pg.SaveToFile(fs, pg.Filename)
	} else if len(pg.Items) == 1 {
		// copy item data up a level
		pg.CopyFromFirstItem()
		filename := pg.Items[0].Filename(prefix)
		return pg.SaveToFile(fs, filename)
	}
	return fmt.Errorf("Nothing to save")
}
