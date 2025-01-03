package internal 

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// fetch all the langs and platforms from the api
func ShowAllLangs()  {
	// API: list all langs
	resp, err := http.Get("https://www.toptal.com/developers/gitignore/api/list")
	if err != nil {
		// TODO: error in api call
		fmt.Println(err.Error())
		os.Exit(1);
	}
	defer resp.Body.Close()
	
	// check resp status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code", resp.StatusCode)
		// TODO: handle better
		os.Exit(1)
	}

	// read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body", err.Error())
		os.Exit(1)
	}


	// parse and print the body in numbered lines
	// read each line
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		words := strings.Split(line, ",")
		// print the first char of each level
		fmt.Printf("%c <-----------------------> \n", words[0][0])
		for _, word := range words {
			fmt.Println(" ", " ", word)
		}
	}
}

func GetTemplate(lang string) string {
	// sanitize input string: remove unwanted spaces
	lang = strings.TrimSpace(lang)

	// only include alphanumeric chars, hypens and underscores
	lang = regexp.MustCompile("[^a-zA-Z0-9_-]").ReplaceAllString(lang, "")	
	if lang == "" {
		fmt.Println("ERROR: language must not be empty")
		os.Exit(1)
	}

	// check if the language exists
	resp, err := http.Get("https://www.toptal.com/developers/gitignore/api/"+ lang)
	if err != nil {
		// TODO: error in api call
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	
	// check resp status code
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("ERROR: programming language `%s` is invalid or not available\n", lang)
		os.Exit(1)
	}
	
	if resp.StatusCode != http.StatusOK {
		fmt.Println("ERROR: Unexpected status code from API")
		os.Exit(1)
	}

	// parse the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR: unable to read response body", err.Error())
		os.Exit(1)
	}

	return (string(body))
}


