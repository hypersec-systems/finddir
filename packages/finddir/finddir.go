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

func Finddir(wordlistPath string, urlbase string) error {
	fmt.Println("Running...")
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
			fmt.Printf("directory found ==> %s | status code %d\n", url, resp.StatusCode)
	
		}
	}
	
	fmt.Println("Done...")
	return nil
}