package utils

import (
	"io"
	"net/http"
	"os"
)

func DownloadFileFromUrl(url, file string) {
	out, err := os.Create(file)
	HandleError(err, "Could not create file to write to.", true)

	resp, err := http.Get(url)
	HandleError(err, "Get request failed. Did API change or network failure or is there service down.", true)
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	HandleError(err, "Could not copy data from request to file.", true)
}
