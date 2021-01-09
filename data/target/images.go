package target

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-windmill/workerpool"
	"github.com/spf13/afero"
)

// DownloadImages is
func (t *Target) DownloadImages(fs afero.Fs, dir string) {

	size := 20
	pool := workerpool.New("imageDownload")
	pool.Start(size)

	items := &t.All
	countM := 0
	countT := 0
	for i := range *items {
		pool.Add(func() {
			e := (*items)[i].DownloadMain(&fs, &dir)
			if e == nil {
				countM++
			} else {
				log.Println((*items)[i].Images.Main)
			}
		})
		pool.Add(func() {
			e := (*items)[i].DownloadThumb(&fs, &dir)
			if e == nil {
				countT++
			} else {
				log.Println((*items)[i].Images.Thumb)
			}
		})
	}
	pool.CloseAndWait()
	l := len(t.All)
	if countM != l || countT != l {
		log.Fatalf("Length did not match downloaded\n%v\n%v\n%v\n", countM, countT, l)
	}
}

// DownloadMain is
func (i *Item) DownloadMain(fs *afero.Fs, dir *string) error {
	file := fmt.Sprintf("%s%s/main/%s.png", *dir, i.Type.Slug, i.Slug)
	e := dl(
		i.Images.Main,
		file,
		*fs)

	if e == nil {
		i.Images.Main = strings.ReplaceAll(file, "./src", "")
	}
	return e
}

// DownloadThumb is
func (i *Item) DownloadThumb(fs *afero.Fs, dir *string) error {
	file := fmt.Sprintf("%s%s/thumb/%s.png", *dir, i.Type.Slug, i.Slug)
	e := dl(
		i.Images.Thumb,
		file,
		*fs)

	if e == nil {
		i.Images.Thumb = strings.ReplaceAll(file, "./src", "")
	}
	return e
}

func dl(url string, filename string, fs afero.Fs) error {
	fmt.Printf("[%v] => %v\n", url, filename)
	resp, gerr := http.Get(url)
	if gerr != nil {
		return gerr
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return e
	}

	fileDir := filepath.Dir(filename)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		fs.MkdirAll(fileDir, os.ModePerm)
	}
	return afero.WriteFile(fs, filename, body, os.FileMode(int(0777)))
}
