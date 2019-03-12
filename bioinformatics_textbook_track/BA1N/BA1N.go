package main

import (
  "os"
  "fmt"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string, int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := strings.Fields(string(b))
  pattern := text[0]
  d, _ := strconv.Atoi(text[1])

  return pattern, d
}


func NextIndex(ix []int, lens int) {
  for j := len(ix) - 1; j >= 0; j-- {
    ix[j]++
    if j == 0 || ix[j] < lens {
      return
    }
    ix[j] = 0
  }
}


func KMerComposition(k int, symbols []string) ([]string) {

  lens := len(symbols)
  kmcs := make([]string, 0)

  for ix := make([]int, k); ix[0] < lens; NextIndex(ix, lens) {
    res := make([]string, k)
    for i, j := range ix {
      res[i] = symbols[j]
    }
    kmcs = append(kmcs, strings.Join(res, ""))
  }

  return kmcs
}


func HammingDistance(seq1, seq2 string) (int) {

  hd := 0
  for i := 0; i < len(seq1); i++ {
    if seq1[i] != seq2[i] {
      hd += 1
    }
  }

  return hd
}


func Neighbors(pattern string, d int) ([]string) {

  kmcs := KMerComposition(len(pattern), []string{"A", "C", "G", "T"})
  neighbors := []string{}
  for _, k := range kmcs {
    if HammingDistance(pattern, k) <= d {
      neighbors = append(neighbors, k)
    }
  }

  return neighbors
}


func Write2File(file_path string, content []string) {
  f, err := os.Create(file_path)
  if err != nil {
      panic(err)
  }
  defer f.Close()

  for _, c := range content {
     fmt.Fprintln(f, c)
  }
}


func main() {
  
  pattern, d := LoadData("rosalind_ba1n.txt")
  
  neighbors := Neighbors(pattern, d)
  Write2File("ba1n_output.txt", neighbors)
}