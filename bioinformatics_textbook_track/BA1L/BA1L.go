package main

import (
  "fmt"
  "strings"
  "math/big"
  "io/ioutil"
)


func LoadData(file_path string) (string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  pattern := strings.TrimSuffix(text, "\r\n")

  return pattern
}


func PatternToNumber(pattern string) (*big.Int) {

  num := big.NewInt(0)
  big4 := big.NewInt(4)
  symbols := map[string]int{"A": 0, "C": 1, "G": 2, "T": 3}

  for _, p := range pattern {
    sp := string(p)
    num = num.Mul(num, big4)
    num = num.Add(num, big.NewInt(int64(symbols[sp])))
  }

  return num
}


func main() {
  
  pattern := LoadData("rosalind_ba1l.txt")
  num := PatternToNumber(pattern)
  fmt.Println(num.String())
}