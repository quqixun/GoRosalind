package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func String2IntSlice(str string) ([]int) {

  ints := []int{}
  int_strs := strings.Fields(str)
  for _, is := range int_strs {
    i, _ := strconv.Atoi(is)
    ints = append(ints, i)
  }

  return ints
}


func LoadData(file_path string) (int, int, []int, []int) {

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

  n, _ := strconv.Atoi(content[0])
  m, _ := strconv.Atoi(content[1])
  A := String2IntSlice(content[2])
  K := String2IntSlice(content[3])

  return n, m, A, K
}


func main() {
  
  n, m, A, K := LoadData("rosalind_bins.txt")
  fmt.Print(n, m, A, K)
}