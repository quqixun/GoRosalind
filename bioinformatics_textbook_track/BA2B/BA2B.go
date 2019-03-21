package main

import (
  "fmt"
  "math"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, []string) {

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

  k, _ := strconv.Atoi(content[0])
  DNAs := content[1:len(content)]

  return k, DNAs
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


func FindMedianString(DNAs []string, k int) (string) {

  kmcs := KMerComposition(k, []string{"A", "C", "G", "T"})

  res := ""
  kmcs_d := make(map[string]int)
  min_sum_d := math.Inf(1)

  for _, kmc := range kmcs {
    for _, DNA := range DNAs {
      min_d := math.Inf(1)
      for i := 0; i <= len(DNA) - k; i++ {
        subDNA := DNA[i:i + k]
        hd := float64(HammingDistance(kmc, subDNA))
        if hd < min_d {
          min_d = hd
        }
      }
      kmcs_d[kmc] += int(min_d)
    }
    if float64(kmcs_d[kmc]) < min_sum_d {
      min_sum_d = float64(kmcs_d[kmc])
      res = kmc
    }
  }

  return res
}


func main() {
  
  k, DNAs := LoadData("rosalind_ba2b.txt")

  res := FindMedianString(DNAs, k)
  fmt.Println(res)
}