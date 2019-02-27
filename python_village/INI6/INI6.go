package main

import (
  "fmt"
  "strings"
  "io/ioutil"
)


func load_words(file_path string) ([]string, bool) {

  words := []string{}

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return words, false
  }

  text := string(b)
  words = strings.Fields(text)

  return words, true
}


func get_keys(dict map[string]int) ([]string) {

  keys := []string{}
  for k, _ := range dict {
    keys = append(keys, k)
  }

  return keys
}


func find_in_slice(word string, keys []string) (bool) {

  is_find := false
  for _, k := range keys {
    if word == k {
      is_find = true
      break
    }
  }

  return is_find
}


func words_dict(words []string) (map[string]int) {

  dict := make(map[string]int)

  for _, word := range words {
    keys := get_keys(dict)
    if find_in_slice(word, keys) {
      dict[word] += 1
    } else {
      dict[word] = 1
    }
  }

  return dict
}


func main() {
  
  words, is_load_success := load_words("rosalind_ini6.txt")

  if is_load_success {
    dict := words_dict(words)
    for k, v := range dict {
      fmt.Println(k, v)
    }
  }
}