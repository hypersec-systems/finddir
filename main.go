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

	if args[0] == "" {
		fmt.Println("\033[1;93mWithout a wordlist!!!")
		fmt.Println("Please set a wordlits path!!!")
		fmt.Println("ex: /usr/share/wordlists/rockyou.txt\033[0;0m")

	} else {
		wordlistPath = args[0]
	}

	if args[1] == "" {
		fmt.Println("\033[1;93mWithout a url!!!")
		fmt.Println("Please set a url")
		fmt.Println("ex: http://exemple.com/\033[0;0m")

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