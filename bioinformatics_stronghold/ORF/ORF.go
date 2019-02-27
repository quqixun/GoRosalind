package main

import (
  "fmt"
  "bufio"
  "strings"
  "io/ioutil"
)


func read_DNA(file_path string) (string, bool) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return "", false
  }

  text := string(b)
  scanner := bufio.NewScanner(strings.NewReader(text))

  DNA := ""
  for scanner.Scan() {
    line_text := scanner.Text()
    if strings.Index(line_text, ">") == -1 {
      DNA += line_text
    }
  }
  return DNA, true
}


func load_table(file_path string) (map[string]string, bool) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return nil, false
  }

  text := string(b)
  scanner := bufio.NewScanner(strings.NewReader(text))

  table := make(map[string]string)
  for scanner.Scan() {
    line_text := scanner.Text()
    line_words := strings.Fields(line_text)
    table[line_words[0]] = line_words[1]
  }

  return table, true
}


func reverse_complement(DNA string) (string) {

  DNA_rc := ""
  for _, s := range DNA {
    switch s {
      case 'A': DNA_rc = "T" + DNA_rc
      case 'T': DNA_rc = "A" + DNA_rc
      case 'C': DNA_rc = "G" + DNA_rc
      case 'G': DNA_rc = "C" + DNA_rc
    }
  }
  return DNA_rc
}


func translate(DNA string, DNA_codon map[string]string) ([]string) {

  ul := 3
  start_idx := []int{}
  proteins := []string{}

  len_DNA := len(DNA)
  for i := 0; i < len_DNA - ul; i++ {
    if DNA_codon[DNA[i:i+ ul]] == "M" {
      start_idx = append(start_idx, i)
    }
  }

  for _, iM := range start_idx {
    protein := ""
    len_sub_DNA := len_DNA - iM
    iter_num := (len_sub_DNA - len_sub_DNA % ul) / ul

    for i := 0; i < iter_num; i++ {
      p := DNA_codon[DNA[iM + ul * i:iM + ul * (i + 1)]]
      if p == "Stop" {
        proteins = append(proteins, protein)
        break
      }
      protein += p
    }
  }

  return proteins
}


func unique(string_slice []string) []string {

  keys := make(map[string]bool)
  list := []string{}

  for _, s := range string_slice {
    if _, value := keys[s]; !value {
      keys[s] = true
      list = append(list, s)
    }
  }    
  return list
}


func main() {

  DNA, is_read_success := read_DNA("rosalind_orf.txt")
  DNA_codon, is_table_success := load_table("DNA_codon_table.txt")

  if is_read_success && is_table_success {
    nc_proteins := translate(DNA, DNA_codon)
    rc_proteins := translate(reverse_complement(DNA), DNA_codon)

    proteins := unique(append(nc_proteins, rc_proteins...))
    proteins_str := strings.Join(proteins, "\n")
    fmt.Println(proteins_str)
  }
}