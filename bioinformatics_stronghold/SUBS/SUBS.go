package main

import (
  "fmt"
  "strconv"
  "strings"
  "io/ioutil"
)

func load_st(file_path string) (string, string, bool) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return "", "", false
  }

  sequences := strings.Fields(string(b))
  s, t := sequences[0], sequences[1]

  return s, t, true
}


func find_motif(s string, t string) ([]int, bool) {

  len_s, len_t := len(s), len(t)

  if len_t > len_s {
    return []int{}, false
  }

  if len_t == len_s {
    if s == t {
      return []int{1}, true
    } else {
      return []int{}, true
    }
  }

  motifs := []int{}
  iter_num := (len_s - len_t) + 1

  for i := 0; i < iter_num; i++ {
    sub_s := s[i:len_t + i]
    if sub_s == t {
      motifs = append(motifs, i + 1)
    }
  }

  return motifs, true
}


func slice2string(motifs []int) (string) {

  motifs_temp := []string{}
  for _, motif := range motifs {
    motifs_temp = append(motifs_temp, strconv.Itoa(motif))
  }

  motifs_str := strings.Join(motifs_temp, " ")
  return motifs_str
}


func main() {

  s, t, is_load_success := load_st("rosalind_subs.txt")

  if is_load_success {
    motifs, is_success := find_motif(s, t)
    if is_success {            
      fmt.Println(slice2string(motifs))
    }
  }
}