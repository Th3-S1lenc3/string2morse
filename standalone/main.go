package main

import (
  "flag"
  "log"
  "fmt"

  s2m "github.com/Th3-S1lenc3/string2morse"
)

func main() {
  convert := s2m.NewString2Morse()

  message := flag.String("m", "", "String to convert.")
  appDir := flag.String("appDir", "", "Path to app config dir, if part of larger application.")

  flag.Parse()

  err = convert.Init(*appDir)
  if err != nil {
    log.Fatal(err)
  }

  morseStr, err := convert.Encode(*message)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(morseStr)
}
