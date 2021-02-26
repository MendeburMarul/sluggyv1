package main

import (
  "bufio"
  "fmt"
  "os"
  "sluggy"
  )

func main() {
  lic := make(chan string)
  lexer := sluggy.NewLexer(lic)

  for{
    r := bufio.NewReader(os.Stdin)
  }
}