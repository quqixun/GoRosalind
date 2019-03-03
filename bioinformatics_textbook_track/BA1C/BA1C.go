package main

import (
  "fmt"
  "io/ioutil"
)


func LoadData(file_path string) (string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  return string(b)
}


func ReverseComplement(DNA string) (string) {

  rc_DNA := ""
  for _, c := range DNA {
    sc := string(c)
    switch sc {
      case "A": rc_DNA = "T" + rc_DNA
      case "T": rc_DNA = "A" + rc_DNA
      case "C": rc_DNA = "G" + rc_DNA
      case "G": rc_DNA = "C" + rc_DNA
    }
  }

  return rc_DNA
}


func main() {
  
  DNA := LoadData("rosalind_ba1c.txt")
  rc_DNA := ReverseComplement(DNA)
  fmt.Println(rc_DNA)
}