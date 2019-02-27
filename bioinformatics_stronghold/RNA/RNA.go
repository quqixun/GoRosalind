package main

import (
  "fmt"
  "strings"
  "io/ioutil"
)


func read_text(file_path string) (string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return "error"
  }

  text := string(b)
  return text
}


func DNA2RNA(DNA string) (string) {

  RNA := strings.Replace(DNA, "T", "U", -1)
  return RNA
}


func main() {

  DNA := read_text("rosalind_rna.txt")

  if DNA != "error" {
    RNA := DNA2RNA(DNA)
    fmt.Println(RNA)
  }
}