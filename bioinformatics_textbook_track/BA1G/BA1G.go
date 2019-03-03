package main

import (
  "fmt"
  "bufio"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string, string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := []string{}
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  DNA1, DNA2 := content[0], content[1]
  return DNA1, DNA2
}


func HammingDistance(DNA1, DNA2 string) (int) {

  hd := 0
  for i := 0; i < len(DNA1); i++ {
    if DNA1[i] != DNA2[i] {
      hd += 1
    }
  }

  return hd
}


func main() {
  
  DNA1, DNA2 := LoadData("rosalind_ba1g.txt")
  hd := HammingDistance(DNA1, DNA2)
  fmt.Println(hd)
}