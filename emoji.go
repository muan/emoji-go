package main

import "os"
import "fmt"

func main()  {
  if len(os.Args) != 2 {
    fmt.Println("Please provide one keyword.")
  } else {
    query := os.Args[1]
    fmt.Println(query)
  }
}
