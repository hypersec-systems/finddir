package main

import (
	"log"
	"find-dir/packages/finddir"
)

func main(){
	wordlistPath := "../find-dir/wordlist/medium.txt"
	err := finddir.Finddir(wordlistPath, "http://google.com/")

	if err != nil {
		log.Panicln(err.Error())
	}
}
