package main

import (
	"finddir/packages/finddir"
	"fmt"
	"log"
	"os"
)

func main(){
	banner()

	var wordlistPath string
	var urlBase string

	args := os.Args[1:]

	if args[0] == "w" {
		wordlistPath = "medium.txt"

	} else {
		wordlistPath = args[0]
	}

	if args[1] == "" {
		fmt.Println("Without a url!!!")
		fmt.Println("ex: http://exemple.com/")

	} else {
		urlBase = args[1]
	}
	
	err := finddir.Finddir(wordlistPath, urlBase)

	if err != nil {
		log.Panicln(err.Error())
	}
}


func banner(){
	fmt.Println("________________________________________")
	fmt.Println(`
    ___________   ______  ____  ________ 
   / ____/  _/ | / / __ \/ __ \/  _/ __ \
  / /_   / //  |/ / / / / / / // // /_/ /
 / __/ _/ // /|  / /_/ / /_/ // // _, _/ 
/_/   /___/_/ |_/_____/_____/___/_/ |_|  
		          BY Gust4vo/nyx
`)
	fmt.Println("________________________________________\n")
}