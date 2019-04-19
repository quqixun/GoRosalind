package main

import (
  "fmt"
  "math"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) ([]string, []float64) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := []string{}
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  n, _ := strconv.Atoi(content[0])

  proteins := []string{}
  for i := 1; i < n + 1; i++ {
    proteins = append(proteins, content[i])
  }

  R := []float64{}
  for i := n + 1; i < len(content); i++ {
    f, _ := strconv.ParseFloat(content[i], 64)
    R = append(R, f)
  }

  return proteins, R
}


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


func ProteinMass(protein string, table map[string]float64) (float64) {

  mass := 0.0

  for _, c := range(protein) {
    mass += table[string(c)]
  }

  return math.Round(mass * 100000) / 100000.0
}



func SpectralConvolution(s1, s2 []float64) ([]float64) {

  sc := []float64{}
  for _, e1 := range s1 {
    for _, e2 := range s2 {
      sc = append(sc, math.Round((e1 - e2) * 100000) / 100000.0)
    }
  }

  return sc
}


func  LargestMultiplicity(conv []float64) (int, float64) {
  
  conv_map := make(map[float64]int)
  lm, x := 0, 0.0

  for _, c := range conv {
    conv_map[c] += 1
    if conv_map[c] > lm {
      lm = conv_map[c]
      x = c
    }
  }

  return lm, x
}


func MatchProteinSpectrum(proteins []string, R []float64, table map[string]float64) (int, string) {

  max_protein := ""
  max_multiplicity := 0

  for _, protein := range proteins {

    sub_proteins := []string{}
    lp := len(protein)
    for i := 0; i < lp; i++ {
      if i != lp - 1 {
        sub_proteins = append(sub_proteins, protein[0:i + 1])
      }
      sub_proteins = append(sub_proteins, protein[lp - i - 1:lp])
    }

    sub_proteins_mass := []float64{}
    for _, sub_protein := range sub_proteins {
      sub_proteins_mass = append(sub_proteins_mass, ProteinMass(sub_protein, table))
    }

    sc := SpectralConvolution(R, sub_proteins_mass)
    lm, _ := LargestMultiplicity(sc)

    if lm >= max_multiplicity {
      max_multiplicity = lm
      max_protein = protein
    }
  }

  return max_multiplicity, max_protein
}


func main() {

  proteins, R := LoadData("rosalind_prsm.txt")
  table := LoadTable("monoisotopic_mass_table.txt")
  
  multiplicity, protein := MatchProteinSpectrum(proteins, R, table)
  fmt.Println(multiplicity, "\n", protein)
}