package main

import (
	"log"

	"github.com/InstagramExtractor/src"
)

func main() {

	url := "https://www.instagram.com/google/"
	x := src.GetInstaData(url)

	log.Println(x)

}
