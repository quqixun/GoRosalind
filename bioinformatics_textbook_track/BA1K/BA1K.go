package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (string, int) {

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
  k, _ := strconv.Atoi(content[1])

  return text, k
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


func Frequency(seq string, kmcs []string) ([]int) {

  freq := make([]int, len(kmcs))
  for i, kmc := range kmcs {
    for j := 0; j <= len(seq) - len(kmc); j++ {
      sub_seq := seq[j:j + len(kmc)]
      if sub_seq == kmc {
        freq[i] += 1
      }
    }
  }

  return freq
}


func IntSliceToString(positions []int) (string) {

  int_strs := []string{}
  for _, p := range positions {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func main() {
  
  seq, k := LoadData("rosalind_ba1k.txt")
  kmcs := KMerComposition(k, []string{"A", "C", "G", "T"})
  freq := Frequency(seq, kmcs)
  fmt.Println(IntSliceToString(freq))
}