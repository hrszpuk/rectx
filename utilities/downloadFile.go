package utilities

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(downloadUrl, path string) int64 {
	file, err := os.Create(path)
	Check(err)
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	response, err := client.Get(downloadUrl)
	Check(err)

	size, err := io.Copy(file, response.Body)
	Check(err)

	file.Close()
	response.Body.Close()

	return size
}
