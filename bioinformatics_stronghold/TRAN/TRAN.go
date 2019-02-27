package main

import (
  "fmt"
  "bufio"
  "strings"
  "io/ioutil"
)


func load_DNAs(file_path string) ([]string, bool) {

  DNAs := []string{}

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return DNAs, false
  }

  text := string(b)

  id := ""
  samples := make(map[string]string)

  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    index := strings.Index(line_text, ">")
    if index > -1 {
      id = line_text[1:]
      samples[id] = ""
    } else {
      if id != "" {
        samples[id] += line_text
      }
    }
  }

  for _, v := range samples {
    DNAs = append(DNAs, v)
  }

  return DNAs, true
}


func is_transition(c0, c1 string) (bool) {

  if (c0 == "A" && c1 == "G") ||
     (c0 == "G" && c1 == "A") ||
     (c0 == "C" && c1 == "T") ||
     (c0 == "T" && c1 == "C") {
    return true
  }

  return false
}


func transition_transversion_ratio(DNAs []string) (float64) {

  length := len(DNAs[0])
  transition := 0.0
  transversion := 0.0

  for i := 0; i < length; i++ {

    c0 := string(DNAs[0][i])
    c1 := string(DNAs[1][i])

    if c0 == c1 {
      continue
    }

    if is_transition(c0, c1) {
      transition += 1.0
    } else {
      transversion += 1.0
    }
  }

  return transition / transversion
}


func main() {
  
  DNAs, is_load_success := load_DNAs("rosalind_tran.txt")

  if is_load_success {
    tt_ratio := transition_transversion_ratio(DNAs)
    fmt.Println(tt_ratio)
  }
}