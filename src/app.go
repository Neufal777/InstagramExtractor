package src

import (
	"io/ioutil"
	"log"
	"net/http"
)

func DownloadPage(PageUrl string) {

	resp, err := http.Get(PageUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	ExtractUrl(string(html[:]))
}

func ExtractUrl(html string) {

	//get video url from html downloaded

}

func SaveVideo(VideoUrl string) {

	log.Println(VideoUrl)
}