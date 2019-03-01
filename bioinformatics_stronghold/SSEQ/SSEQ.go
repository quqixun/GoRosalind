package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func load_DNAs(file_path string) (string, string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := make(map[string]string)
  id := ""

  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    if strings.Index(line_text, ">") > -1 {
      id = line_text[1:len(line_text)]
      content[id] = ""
    } else {
      if id != "" {
        content[id] += line_text
      }
    }
  }

  DNAs := []string{}
  for _, v := range content {
    DNAs = append(DNAs, v)
  }

  s, t := "", ""
  if len(DNAs[0]) > len(DNAs[1]) {
    s, t = DNAs[0], DNAs[1]
  } else {
    s, t = DNAs[1], DNAs[0]
  }

  return s, t
}


func find_spliced_motif(s, t string) ([]int) {

  indices := []int{}

  s_idx := -1
  for _, tc := range t {
    tc_idx := strings.Index(s[s_idx + 1:len(s)], string(tc))
    indices = append(indices, tc_idx + 1 + s_idx + 1)
    s_idx += tc_idx + 1
  }

  return indices
}


func main() {

  s, t := load_DNAs("rosalind_sseq.txt")

  indices := find_spliced_motif(s, t)
  
  indices_str := []string{}
  for _, i := range indices {
    i_str := strconv.Itoa(i)
    indices_str = append(indices_str, i_str)
  }

  f_str := strings.Join(indices_str, " ")
  fmt.Println(f_str)
}