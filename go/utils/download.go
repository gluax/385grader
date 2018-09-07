package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFileFromUrl(url, file string) {
	out, err := os.Create(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
