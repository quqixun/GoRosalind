package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string, int, int) {

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

  text := content[0]
  num_strs := strings.Fields(content[1])
  k, _ := strconv.Atoi(num_strs[0])
  d, _ := strconv.Atoi(num_strs[1])

  return text, k, d
}


func FindKMers(text string, k int) (map[string]int) {

  kmers := make(map[string]int)
  for i := 0; i <= len(text) - k; i++ {
    kmers[text[i:i + k]] += 1
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


func MostFrequentMismatches(kmers map[string]int, kmcs []string, d int) ([]string) {
  mismatches := make(map[string]int)
  max_n := 0
  for seq1, n := range kmers {
    for _, seq2 := range kmcs {
      hd := HammingDistance(seq1, seq2)
      if hd <= d {
        mismatches[seq2] += n
        if mismatches[seq2] > max_n {
          max_n = mismatches[seq2]
        }
      }
    }
  }

  res := []string{}
  for k, v := range mismatches {
    if v == max_n {
      res = append(res, k)
    }
  }

  return res
}


func main() {
  
  text, k, d := LoadData("rosalind_ba1i.txt")

  kmers := FindKMers(text, k)
  kmcs := KMerComposition(k, []string{"A", "C", "G", "T"})
  res := MostFrequentMismatches(kmers, kmcs, d)
  fmt.Println(strings.Join(res, " "))
}