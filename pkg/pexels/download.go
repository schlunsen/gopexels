package pexels

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"path"
)

func (s *PexelAPI) DownloadImage(URL string, basePath string) bool {

	r, _ := http.NewRequest("GET", URL, nil)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return false
	}

	// open the file

	file, err := os.Create(basePath + path.Base(r.URL.Path))

	if err != nil {
		log.Println(err)
		return false
	}

	defer file.Close()

	// create the reader and writters we'll connect
	respReader := bufio.NewReader(resp.Body)
	fileWriter := bufio.NewWriter(file)

	// write the content of the resp to the file
	_, err = respReader.WriteTo(fileWriter)

	if err != nil {
		return false
	}

	return true
}
