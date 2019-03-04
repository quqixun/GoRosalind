package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string, string, int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  t := string(b)
  content := []string{}
  scanner := bufio.NewScanner(strings.NewReader(t))
  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  text := content[1]
  pattern := content[0]
  d, _ := strconv.Atoi(content[2])

  return text, pattern, d
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


func ApproximateMatching(text, pattern string, d int) ([]int) {

  idx := []int{}
  for i := 0; i <= len(text) - len(pattern); i++ {
    subtext := text[i:i + len(pattern)]
    if HammingDistance(subtext, pattern) <= d {
      idx = append(idx, i)
    }
  }

  return idx
}


func IntSliceToString(int_slice []int) (string) {

  int_strs := []string{}
  for _, i := range int_slice {
    int_strs = append(int_strs, strconv.Itoa(i))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func main() {
  
  text, pattern, d := LoadData("rosalind_ba1h.txt")
  idx := ApproximateMatching(text, pattern, d)
  fmt.Println(IntSliceToString(idx))
}