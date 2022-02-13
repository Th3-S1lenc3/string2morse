package main

type Signals struct {
  Characters []Signal `json:"characters"`
}

type Signal struct {
  Character string `json:"character"`
  Signal string `json:"signal"`
}
