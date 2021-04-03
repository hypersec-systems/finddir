package finddir

import (
	"os"
	"bufio"
	"fmt"
	"net/http"
)


func readWordlist(path string) ([]string, error) {
	wordlistFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer wordlistFile.Close()

	var wordlistLines []string
	scanner := bufio.NewScanner(wordlistFile)

	for scanner.Scan() {
		wordlistLines = append(wordlistLines, scanner.Text())
	}

	return wordlistLines, scanner.Err()	
}


func writeFile(urls []string) error {
	file, err := os.Create("urlsFound.txt")
	
	if err != nil {
		return err
	}

	defer file.Close()

	for _, url := range(urls){
		_, err := file.WriteString(url + "\n")
	
		if err != nil {
			return err
		}
	}

	return nil
}


func Finddir(wordlistPath string, urlbase string) error {
	fmt.Println("\u001b[32;1mRunning...\u001b[0m")

	var urls []string
	wordlist, err := readWordlist(wordlistPath)

	if err != nil {
		return err
	}

	for _, item := range(wordlist){
		url := urlbase + item
		resp, err := http.Get(url)

		if err != nil {
			return err
		}

		if resp.StatusCode == 200 {
			fmt.Printf("\u001b[36;1m[*] Directory found ==> %s status code %d\u001b[0m\n", url, resp.StatusCode)
			urls = append(urls, url) 
		}
	}
	

	err = writeFile(urls)

	if err != nil {
		return err
	}
	
	fmt.Println("\u001b[32;1mDone...\u001b[0m")
	return nil
}

