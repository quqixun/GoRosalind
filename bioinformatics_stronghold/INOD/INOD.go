package main

import (
  "fmt"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := strings.TrimSuffix(string(b), "\r\n")
  n, _ := strconv.Atoi(string(text))
  return n
}


func InternalNodesNum(n_leafs int) (int) {
  return n_leafs - 2
}


func main() {
  
  n_leafs := LoadData("rosalind_inod.txt")
  n_inode := InternalNodesNum(n_leafs)
  fmt.Println(n_inode)
}