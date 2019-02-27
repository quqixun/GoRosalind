package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func load_content(file_path string) (string, [][]int, bool) {

  s := ""
  indices := [][]int{}

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return s, indices, false
  }

  content := []string{}
  text := string(b)
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    content = append(content, line_text)
  }

  s = content[0]
  numbers := strings.Fields(content[1])

  ints := []int{}
  for _, n := range numbers {
    ni, _ := strconv.Atoi(n)
    ints = append(ints, ni)
  }

  indices = append(indices, []int{ints[0], ints[1]})
  indices = append(indices, []int{ints[2], ints[3]})

  return s, indices, true
}


func find_words(s string, indices [][]int) ([]string) {

  words := []string{}

  for _, idx := range indices {
    words = append(words, s[idx[0]:idx[1] + 1])
  }

  return words
}


func main() {

  s, indices, is_load_success := load_content("rosalind_ini3.txt")

  if is_load_success {
    words := find_words(s, indices)
    words_string := strings.Join(words, " ")
    fmt.Println(words_string)
  }
}