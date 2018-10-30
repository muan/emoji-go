package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	s "strings"
)

type emoji struct {
	Keywords         []string
	Char             string
	FitzpatrickScale bool
	Category         string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide one keyword.")
	} else {
		query := os.Args[1]
		data := getLib()
		fmt.Println("Keyword:", query)
		var results []string
		for key, item := range data {
			if s.Index(key, query) >= 0 {
				if key == query {
					results = append([]string{item.Char}, results...)
				} else {
					results = append(results, item.Char)
				}
			}
		}

		if len(results) > 0 {
			fmt.Println("Found:")
			for i, emoji := range results {
				if i < 5 {
					fmt.Println(i+1, emoji)
				} else {
					break
				}
			}
		} else {
			fmt.Println("No emoji found 😭.")
		}
	}
}

func getLib() map[string]emoji {
	lib, err := ioutil.ReadFile("./emoji.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No emoji.json found. Fetching..")
			lib = fetchLib()
			ioutil.WriteFile("emoji.json", lib, 0777)
		} else {
			fmt.Println(err)
		}
	}

	var data map[string]emoji
	json.Unmarshal([]byte(string(lib)), &data)
	return data
}

func fetchLib() []byte {
	resp, err := http.Get("https://raw.githubusercontent.com/muan/emojilib/master/emojis.json")
	if err != nil {
		fmt.Println("ERRRRR")
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
