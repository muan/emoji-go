package main

import "os"
import "fmt"
import "encoding/json"
import "io/ioutil"
import s "strings"

type emojiT struct {
  Name string
  Keywords []string
  Char string
  Fitzpatrick_scale bool
  Category string
}

func main()  {
  if len(os.Args) != 2 {
    fmt.Println("Please provide one keyword.")
  } else {
    query := os.Args[1]
    data := getLib()
    fmt.Println("Keyword:", query)
    var results []string
    for _, emoji := range data {
      if s.Index(emoji.Name, query) >= 0 {
        if emoji.Name == query {
          results = append([]string{emoji.Char}, results...)
        } else {
          results = append(results, emoji.Char)
        }
      }
    }

    if len(results) > 0 {
      fmt.Println("Found:")
      for i, emoji := range results {
        if i < 5 {
          fmt.Println(i + 1, emoji)
        } else {
          break
        }
      }
    } else {
      fmt.Println("No emoji found 😭.")
    }
  }
}

func getLib() []emojiT {
  lib, _ := ioutil.ReadFile("./emoji-for-go.json")
  var data []emojiT
  json.Unmarshal([]byte(string(lib)), &data)
  return data
}
