package main

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
  "strings"
)

type Convert struct {
  str string
  morseCodeStr string
  dictionary Signals
}

func NewConvert() (*Convert, error) {
  return &Convert{}, nil
}

func (c *Convert) GetDictionary() Signals {
  return c.dictionary
}

func (c *Convert) GetMorseCode() string {
  return c.morseCodeStr
}

func (c *Convert) Encode() (string, error) {
  strArr := strings.Split(c.str, "")

  for i := 0; i < len(strArr); i++ {
    signal, err := c.getSignalForCharacter(strArr[i])
    if err != nil {
      return "", err
    }

    c.morseCodeStr += signal
    if signal != "/" {
      c.morseCodeStr += "/"
    }
  }

  return c.morseCodeStr, nil
}


func (c *Convert) getSignalForCharacter(character string) (string, error) {
  if character == "" {
    return "", fmt.Errorf("Invalid character: \"%s\"", character)
  }

  if character == " " {
    return "/", nil
  }

  characters := c.dictionary.Characters

  for i := 0; i < len(characters); i++ {
    signal := characters[i]

    if signal.Character == character {
      return signal.Signal, nil
    }
  }

  return "", fmt.Errorf("Not Found.")
}

func (c *Convert) Init(str string) error {
  if str == "" {
    return fmt.Errorf("No string provided.")
  }

  c.str = strings.ToLower(str)
  c.morseCodeStr = ""

  jsonData, err := ioutil.ReadFile("signals.json")
  if err != nil {
    return err
  }

  err = json.Unmarshal(jsonData, &c.dictionary)
  if err != nil {
    return err
  }

  return nil
}
