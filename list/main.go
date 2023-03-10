package main

import (
  "io/fs"
  "os"
  "fmt"
)

func main() {
  pattern := "*"

  names, err := fs.Glob(os.DirFS("."), pattern)
  if err != nil {
    fmt.Println("Error listing files.", err)
  }

  for _, name := range names {
    fmt.Printf("%v\n", name)
  }
}

