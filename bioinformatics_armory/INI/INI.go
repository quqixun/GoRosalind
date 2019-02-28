package main

import (
  "fmt"
  "strconv"
  "io/ioutil"
)


func load_DNA(file_path string) (string, bool) {

  DNA := ""

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return DNA, false
  }

  DNA = string(b)
  return DNA, true
}


func count_symbols(DNA string) (map[string]int) {

  counts := make(map[string]int)
  counts["A"] = 0
  counts["C"] = 0
  counts["G"] = 0
  counts["T"] = 0

  for _, c := range DNA {
    cs := string(c)
    switch cs {
      case "A": counts["A"] += 1
      case "C": counts["C"] += 1
      case "G": counts["G"] += 1
      case "T": counts["T"] += 1
    }
  }

  return counts
}


func main() {
    
  DNA, is_load_success := load_DNA("rosalind_ini.txt")

  if is_load_success {
    counts := count_symbols(DNA)
    counts_str := strconv.Itoa(counts["A"]) +
                  strconv.Itoa(counts["C"]) +
                  strconv.Itoa(counts["G"]) +
                  strconv.Itoa(counts["T"])
    fmt.Println(counts_str)
  }
}