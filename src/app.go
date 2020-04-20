package src

import (
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"regexp"
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

	fmt.Println(string(html[:]))
	//ExtractUrl(string(html[:]))
}

func ExtractUrl(html string) {

	//get video url from html downloaded
	re := regexp.MustCompile(`"contentUrl":"[^"]+`)
	res := re.FindString(html)
	 
	 
	 log.Println(res)

}

func SaveVideo(VideoUrl string) {

	log.Println(VideoUrl)
}
