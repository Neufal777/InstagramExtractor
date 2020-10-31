package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type User struct {
	Id           string
	Url          string
	Username     string
	Description  string
	ProfileImage string
}

/*
	Given a profile, it downloads its HTML
*/
func DownloadPage(PageUrl string) {

	fmt.Println("Downloading profile..")
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

	//sending the html to be processed
	GetInstaData(h)

}

func GetInstaData(content string) {

	/*
		We process the content (html of the website) to extract information
	*/
	var (

		//regex for data extraction
		getProfileImage = regexp.MustCompile("property=.og:image. content=.([^\"']*)")                                   //Gets the profile image
		getUsername     = regexp.MustCompile("<meta property=.al:ios:url. content=.instagram://user.username=([^\"']*)") //Gets the username
		getId           = regexp.MustCompile("<meta property=.al:ios:app_store_id. content=.([^\"']*). />")              //Gets the Id of the user
		getDescription  = regexp.MustCompile("{.user.:{.biography.:.([^\"']*)")                                          //Gets the description
		getUrl          = regexp.MustCompile("<meta property=.og:url. content=.([^\"']*). />")                           //Gets the url of the profile

	)

	//filling the user struct with the data extracted
	user := User{
		Id:           getId.FindStringSubmatch(content)[1],
		Url:          getUrl.FindStringSubmatch(content)[1],
		Username:     getUsername.FindStringSubmatch(content)[1],
		Description:  getDescription.FindStringSubmatch(content)[1],
		ProfileImage: getProfileImage.FindStringSubmatch(content)[1],
	}

	fmt.Printf("%+v\n", user)

}
