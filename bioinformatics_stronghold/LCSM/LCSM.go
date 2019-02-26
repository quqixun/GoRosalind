package main

import (
  "fmt"
  "math"
  "bufio"
  "strings"
  "io/ioutil"
)


func load_text(file_path string) (string, bool) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return "", false
  }

  text := string(b)
  return text, true
}


func load_fast(text string) (map[string]string) {

  samples := make(map[string]string)

  id := ""
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    i := strings.Index(line_text, ">")
    if i > -1 {
      id = line_text[1:]
      samples[id] = ""
    } else {
      if id != "" {
        samples[id] += line_text
      }
    }
  }

  return samples
}


func find_shortest(sentences []string) (int) {

  shortest_len := math.Inf(1)
  shortest_index := -1

  for i, s := range sentences {
    len_s := float64(len(s))
    if len_s < shortest_len {
      shortest_len = len_s
      shortest_index = i
    }
  }

  return shortest_index
}


func find_shared_motif(samples map[string]string) (string) {

  sentences := []string{}
  for _, v := range samples {
    sentences = append(sentences, v)
  }
  n_sentences := len(sentences)
  
  shortest_idx := find_shortest(sentences)
  shortest_string := sentences[shortest_idx]
  shortest_len := len(shortest_string)

  done := false
  fsm := ""
  for l := shortest_len; l > 0; l-- {
    for i := 0; i <= shortest_len - l; i++ {
      substring := shortest_string[i:i + l]
      n_ss := 0
      for _, sentence := range sentences {
        if strings.Index(sentence, substring) > -1 {
          n_ss += 1
        }
      }
      if n_ss == n_sentences {
        fsm = substring
        done = true
        break
      }
    }
    if done {
      break
    }
  }

  return fsm
}


func main() {

  text, is_load_success := load_text("rosalind_lcsm.txt")
  if is_load_success {
    samples := load_fast(text)
    fsm := find_shared_motif(samples)
    fmt.Println(fsm)
  }
}