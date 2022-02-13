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
  appDir := flag.String("appDir", "", "Path to app config dir, if part of larger application.")

  flag.Parse()

  err = convert.Init(*message, *appDir)
  if err != nil {
    log.Fatal(err)
  }

  morseStr, err := convert.Encode()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(morseStr)
}
