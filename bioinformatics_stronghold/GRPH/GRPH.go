package main

import (
  "fmt"
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


func load_samples(text string) (map[string]string, []string) {

  samples := make(map[string]string)
  ordered_keys := []string{}

  id := ""
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    i := strings.Index(line_text, ">")
    if i > -1 {
      id = line_text[1:]
      samples[id] = ""
      ordered_keys = append(ordered_keys, id)
    } else {
      if id != "" {
        samples[id] += line_text
      }
    }
  }

  return samples, ordered_keys
}


func overlap_graphs(samples map[string]string, ordered_keys []string) ([]string) {

  values := []string{}
  for _, key := range ordered_keys {
    values = append(values, samples[key])
  }

  k := 3
  graphs := []string{}
  for i := 0; i < len(ordered_keys); i++ {
    DNA1 := values[i]
    DNA1_len := len(DNA1)
    for j := i + 1; j < len(ordered_keys); j++ {
      DNA2 := values[j]
      DNA2_len := len(DNA2)

      if DNA1[DNA1_len - k:] == DNA2[0:k] {
        graphs = append(graphs, ordered_keys[i] + " " + ordered_keys[j])
      }

      if DNA2[DNA2_len - k:] == DNA1[0:k] {
        graphs = append(graphs, ordered_keys[j] + " " + ordered_keys[i])
      }
    }
  }

  return graphs
}


func main() {

  text, is_text_success := load_text("rosalind_grph.txt")

  if is_text_success {
    samples, ordered_keys := load_samples(text)
    graphs := overlap_graphs(samples, ordered_keys)
    for _, graph := range graphs {            
      fmt.Println(graph)
    }
  }
}