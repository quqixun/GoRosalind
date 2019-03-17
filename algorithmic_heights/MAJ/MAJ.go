package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, int, [][]int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := [][]int{}
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    int_strs := strings.Fields(scanner.Text())
    ints := []int{}
    for _, str := range int_strs {
      i, _ := strconv.Atoi(str)
      ints = append(ints, i)
    }
    content = append(content, ints)
  }

  k := content[0][0]
  n := content[0][1]
  arrays := content[1:len(content)]
  return k, n, arrays
}


func FindMajority(array []int) (int) {

  half := len(array) / 2
  element_lens := make(map[int]int)

  for _, v := range array {
    element_lens[v] += 1
    if element_lens[v] > half {
      return v
    }
  }

  return -1
}


func IntSliceToString(ints []int) (string) {

  int_strs := []string{}
  for _, p := range ints {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func main() {
  
  _, _, arrays := LoadData("rosalind_maj.txt")

  majorities := []int{}
  for _, array := range arrays {
    majorities = append(majorities, FindMajority(array))
  }
  fmt.Println(IntSliceToString(majorities))
}