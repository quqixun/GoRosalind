package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  DNA := ""
  text := string(b)
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    if strings.Index(line_text, ">") == -1 {
      DNA += line_text
    }
  }

  return DNA
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


func CountKMers(DNA string, kmcs []string) ([]int) {

  count := []int{}
  for _, kmc := range kmcs {
    n_kmc := 0
    for i := 0; i <= len(DNA) - len(kmc); i++ {
      subDNA := DNA[i:i + len(kmc)]
      if subDNA == kmc {
        n_kmc += 1
      }
    }
    count = append(count, n_kmc)
  }

  return count
}


func IntSliceToString(int_slice []int) (string) {

  int_strs := []string{}
  for _, i := range int_slice {
    int_strs = append(int_strs, strconv.Itoa(i))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func main() {
  
  DNA := LoadData("rosalind_kmer.txt")

  k := 4
  kmers := KMerComposition(k, []string{"A", "C", "G", "T"})
  count := CountKMers(DNA, kmers)
  fmt.Println(IntSliceToString(count))
}