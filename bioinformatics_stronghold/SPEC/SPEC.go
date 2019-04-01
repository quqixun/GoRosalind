package main

import (
  "fmt"
  "math"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadTable(file_path string) (map[string]float64) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  table := make(map[string]float64)
  scanner := bufio.NewScanner(strings.NewReader(string(b)))
  for scanner.Scan() {
    line_text := scanner.Text()
    words := strings.Fields(line_text)
    table[words[0]], _ = strconv.ParseFloat(words[1], 64)
  }

  return table
}


func LoadData(file_path string) ([]float64) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  L := []float64{}
  scanner := bufio.NewScanner(strings.NewReader(string(b)))
  for scanner.Scan() {
    line_text := scanner.Text()
    n, _ := strconv.ParseFloat(line_text, 64)
    L = append(L, n)
  }

  return L
}


func Keep4Digits(v float64) (float64) {
  return math.Round(v * 10000) / 10000.0
}


func FindAminoAcidByWeight(mass_table map[string]float64, weight float64) (string) {

  for k, v := range mass_table {
    rv := Keep4Digits(v)
    if rv == weight {
      return k
    }
  }

  return ""
}


func ProteinWithPrefixSpectrum(L []float64, mass_table map[string]float64) (string) {

  protein := ""
  for i := 1; i < len(L); i++ {
    weight := L[i] - L[i - 1]
    weight = Keep4Digits(weight)
    protein += FindAminoAcidByWeight(mass_table, weight)
  }

  return protein
}


func main() {
  mass_table := LoadTable("monoisotopic_mass_table.txt")
  L := LoadData("rosalind_spec.txt")
  protein := ProteinWithPrefixSpectrum(L, mass_table)
  fmt.Println(protein)
}