package target

import (
	"acnh/utility"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-windmill/workerpool"
	"github.com/spf13/afero"
)

// DownloadImages is
func (t *Target) DownloadImages(fs afero.Fs, dir string) {

	size := 5
	pool := workerpool.New("imageDownload")
	pool.Start(size)

	for x := range t.All {
		i := &(t.All)[x]
		pool.Add(func() {
			str := strings.TrimRight(strings.ReplaceAll(dir, "./src/", ""), "/")
			i.Images.Main = utility.URL(str, i.Slug, "main") + ".png"
			dl(i.Images.MainSource, "./src"+i.Images.Main, fs)
		})
		pool.Add(func() {
			str := strings.TrimRight(strings.ReplaceAll(dir, "./src/", ""), "/")
			i.Images.Thumb = utility.URL(str, i.Slug, "thumb") + ".png"
			dl(i.Images.ThumbSource, "./src"+i.Images.Thumb, fs)
		})
	}
	pool.CloseAndWait()

}

func dl(url string, filename string, fs afero.Fs) (string, error) {
	fmt.Printf("[%v] => %v\n", url, filename)
	resp, gerr := http.Get(url)
	if gerr != nil {
		return "", gerr
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return "", e
	}

	fileDir := filepath.Dir(filename)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		fs.MkdirAll(fileDir, os.ModePerm)
	}
	return filename, afero.WriteFile(fs, filename, body, os.FileMode(int(0777)))
}
