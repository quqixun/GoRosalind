package main

import (
  "fmt"
  "bufio"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) ([]string) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  id := ""
  text := string(b)
  proteins := []string{}
  content := make(map[string]string)
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    if strings.Index(line_text, ">") > -1 {
      id = line_text[1:len(line_text)]
      proteins = append(proteins, id)
      content[id] = ""
    } else {
      if id != "" {
        content[id] += line_text
      }
    }
  }

  seqs := []string{}
  for _, p := range proteins {
    seqs = append(seqs, content[p])
  }

  return seqs
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


func DistanceMatrix(seqs []string) ([][]float64) {

  dm := [][]float64{}
  for _, s1 := range(seqs) {
    sdm := []float64{}
    for _, s2 := range(seqs) {
      pd := float64(HammingDistance(s1, s2)) / float64(len(s1))
      sdm = append(sdm, pd)
    }
    dm = append(dm, sdm)
  }

  return dm
}


func Mat2String(mat [][]float64) (string) {

  ms := []string{}
  for _, row := range mat {
    rs := []string{}
    for _, item := range row {
      item_str := fmt.Sprintf("%.6f", item)
      rs = append(rs, item_str)
    }
    row_str := strings.Join(rs, " ")
    ms = append(ms, row_str)
  }

  mat_str := strings.Join(ms, "\r\n")
  return mat_str
}


func main() {
  
  seqs := LoadData("rosalind_pdst.txt")
  dm := DistanceMatrix(seqs)
  ms := Mat2String(dm)
  fmt.Println(ms)
}