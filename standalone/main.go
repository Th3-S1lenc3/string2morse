package main

import (
  "flag"
  "log"
  "fmt"

  s2m "github.com/Th3-S1lenc3/string2morse"
)

func main() {
  convert, err := s2m.NewConvert()
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
