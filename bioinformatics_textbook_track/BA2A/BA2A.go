package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, int, []string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  tb := string(b)
  content := []string{}
  scanner := bufio.NewScanner(strings.NewReader(tb))
  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  num_strs := strings.Fields(content[0])
  k, _ := strconv.Atoi(num_strs[0])
  d, _ := strconv.Atoi(num_strs[1])
  text := content[1:len(content)]

  return k, d, text
}


func FindKMers(text []string, k int) (map[string]int) {

  kmers := make(map[string]int)
  for _, t := range text {
    for i := 0; i <= len(t) - k; i++ {
      kmers[t[i:i + k]] += 1
    }
  }
  
  return kmers
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


func InSlice(t string, strs []string) (bool) {

  for _, s := range strs {
    if t == s {
      return true
    }
  }

  return false
}


func MotifEnumeration(DNAs []string, k, d int) ([]string) {

  kmcs := KMerComposition(k, []string{"A", "C", "G", "T"})
  kmers := make(map[string]int)
  for _, DNA := range DNAs {
    for _, kmc := range kmcs {
      appear := false
      for i := 0; i <= len(DNA) - k; i++ {
        subDNA := DNA[i:i + k]
        if HammingDistance(subDNA, kmc) <= d {
          appear = true
          break
        }
      }
      if appear {
        kmers[kmc] += 1
      }
    }
  }

  patterns := []string{}
  for key, value := range kmers {
    if value == len(DNAs) {
      if !InSlice(key, patterns) {
        patterns = append(patterns, key)
      }
    }
  }

  return patterns
}


func main() {
  
  k, d, DNAs := LoadData("rosalind_ba2a.txt")

  res := MotifEnumeration(DNAs, k, d)
  fmt.Println(strings.Join(res, " "))
}