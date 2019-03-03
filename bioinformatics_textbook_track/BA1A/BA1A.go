package main

import (
  "fmt"
  "bufio"
  "strings"
  "io/ioutil"
)


func load_data(file_path string) (string, string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  bs := string(b)
  content := []string{}
  scanner := bufio.NewScanner(strings.NewReader(bs))
  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  text, pattern := content[0], content[1]
  return text, pattern
}


func PatternCount(text, pattern string) (int) {

  count := 0
  for i := 0; i <= len(text) - len(pattern); i++ {
    subtext := text[i:i + len(pattern)]
    if subtext == pattern {
      count += 1
    }
  }

  return count
}


func main() {
  
  text, pattern := load_data("rosalind_ba1a.txt")
  count := PatternCount(text, pattern)
  fmt.Println(count)
}