package main

import (
  "fmt"
  "math"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  int_strs := strings.Fields(text)

  index, _ := strconv.Atoi(int_strs[0])
  k, _ := strconv.Atoi(int_strs[1])
  return index, k
}


func NumbertoPattern(index, k int) (string) {

  symbols := map[int]string{
    0: "A", 1: "C", 2: "G", 3: "T",
  }

  if k == 1 {
    return symbols[index]
  }

  mod := int(math.Mod(float64(index), 4.0))
  index = int(math.Floor(float64(index) / 4.0))
  return NumbertoPattern(index, k - 1) + symbols[mod]
}


func main() {

  index, k := LoadData("rosalind_ba1m.txt")
  pattern := NumbertoPattern(index, k)
  fmt.Println(pattern)
}