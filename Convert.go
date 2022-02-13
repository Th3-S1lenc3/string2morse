package main

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
  "strings"
  "os"
  "time"

  "github.com/cavaliergopher/grab/v3"
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

func (c *Convert) DownloadSignals(configDir string, fileName string) error {
  fmt.Printf("Cannot find \"%s\" in \"%s\"\n", fileName, configDir)

  remoteFileURL := "https://raw.githubusercontent.com/Th3-S1lenc3/string2morse/master/json/signals.min.json"

  // Create Client
  client := grab.NewClient()
  req, _ := grab.NewRequest(configDir, remoteFileURL)

  // Start Download
  fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

  // Start UI Loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
  for {
    select {
    case <-t.C:
      fmt.Printf(
        "  transferred %v / %v bytes (%.2f%%)\n",
        resp.BytesComplete(),
        resp.Size,
        100 * resp.Progress(),
      )
    case <- resp.Done:
      break Loop
    }
  }

  if err := resp.Err(); err != nil {
    return fmt.Errorf("Download failed %v\n", err)
  }

  fmt.Printf("Download saved to %v \n", resp.Filename)

  return nil
}

func (c *Convert) Init(str string, appDir string) error {
  if str == "" {
    return fmt.Errorf("No string provided.")
  }

  c.str = strings.ToLower(str)
  c.morseCodeStr = ""

  if appDir == "" {
    cwd, err := os.UserConfigDir()
    if err != nil {
      return err
    }

    appDir = cwd
  }

  _, err := os.Stat(appDir)
	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("Cannot find directory: \"%s\"", appDir)
	}

  configDir := fmt.Sprintf("%s/String2Morse", appDir)

  _, err = os.Stat(configDir)
  if err != nil && os.IsNotExist(err) {
    err = os.Mkdir(configDir, os.FileMode(0755))
    if err != nil {
      return err
    }
	}

  signalsJsonFilePath := fmt.Sprintf("%s/signals.min.json", configDir)

  _, err = os.Stat(signalsJsonFilePath)
  if err != nil && os.IsNotExist(err) {
    err = c.DownloadSignals(configDir, "signals.min.json")
    if err != nil {
      return err
    }
  }

  jsonData, err := ioutil.ReadFile(signalsJsonFilePath)
  if err != nil {
    return err
  }

  err = json.Unmarshal(jsonData, &c.dictionary)
  if err != nil {
    return err
  }

  return nil
}
