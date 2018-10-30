package main

import "os"
import "fmt"
import "encoding/json"
import "io/ioutil"
import s "strings"

type emoji struct {
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

func getLib() map[string]emoji {
  lib, _ := ioutil.ReadFile("./emoji.json")
  var data map[string]emoji
  json.Unmarshal([]byte(string(lib)), &data)
  return data
}
