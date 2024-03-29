package utilities

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(downloadUrl, path string) int64 {
	file, err := os.Create(path)
	Check(err, true, "Attempted to create a file to copy bytes to during a download but failed!")
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	response, err := client.Get(downloadUrl)
	Check(err, true, "Attempted to fetch downloadable content but failed!")

	size, err := io.Copy(file, response.Body)
	Check(err, true, "Attmpted to copy bytes from response body to file but failed!")

	file.Close()
	response.Body.Close()

	return size
}
