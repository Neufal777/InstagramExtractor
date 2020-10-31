package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Profile struct {
	Id           string
	Url          string
	Username     string
	Description  string
	ProfileImage string
}

func DownloadPage(PageUrl string) {

	fmt.Println("Downloading..")

	resp, err := http.Get(PageUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	h := string(html[:])

	GetInstaData(h)

}

func GetInstaData(content string) {

	/*
		We process the content (html of the website) to extract url of the video
	*/
	var (
		getProfileImage = regexp.MustCompile("property=.og:image. content=.([^\"']*)")                                   //Gets the user profile image
		getUsername     = regexp.MustCompile("<meta property=.al:ios:url. content=.instagram://user.username=([^\"']*)") //Gets the username
		getId           = regexp.MustCompile("<meta property=.al:ios:app_store_id. content=.([^\"']*). />")              //Gets the username
		getDescription  = regexp.MustCompile("{.user.:{.biography.:.([^\"']*)")                                          //Gets the username
		getUrl          = regexp.MustCompile("<meta property=.og:url. content=.([^\"']*). />")                           //Gets the username

	)

	images := []string{}
	result := getProfileImage.FindStringSubmatch(content)[1]
	images = append(images, result)

	user := Profile{
		Id:           getId.FindStringSubmatch(content)[1],
		Url:          getUrl.FindStringSubmatch(content)[1],
		Username:     getUsername.FindStringSubmatch(content)[1],
		Description:  getDescription.FindStringSubmatch(content)[1],
		ProfileImage: getProfileImage.FindStringSubmatch(content)[1],
	}

	fmt.Println(user)

}
