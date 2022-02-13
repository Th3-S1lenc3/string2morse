package main

import (
  "flag"
  "log"
  "fmt"
)

func main() {
  convert, err := NewConvert()
  if err != nil {
    log.Fatal(err)
  }

  message := flag.String("m", "", "String to convert.")

  flag.Parse()

  err = convert.Init(*message)
  if err != nil {
    log.Fatal(err)
  }

  morseStr, err := convert.Encode()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(morseStr)
}
