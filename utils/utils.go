package utils

import (
	"io/ioutil"
	"net/http"
)

func DownloadHtml(url string) string {

	/*
		This function downloads an html source code and
		turns it into a string
	*/

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	htmlString := string(html[:])

	return htmlString
}
