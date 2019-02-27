package main

import (
  "fmt"
  "bufio"
  "strings"
  "io/ioutil"
)


func load_text(file_path string) (string, bool) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return "", false
  }

  text := string(b)
  return text, true
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


func load_samples(text string) (map[string]string) {

  samples := make(map[string]string)

  id := ""
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    i := strings.Index(line_text, ">")
    if i > -1 {
      id = line_text[1:]
      samples[id] = ""
    } else {
      if id != "" {
        samples[id] += line_text
      }
    }
  }

  return samples
}


func select_DNA_introns(DNAs map[string]string) (string, []string) {

  DNA := ""
  introns := []string{}

  max_len := 0
  for _, v := range DNAs {
    if len(v) > max_len {
      max_len = len(v)

      if DNA != "" {
        introns = append(introns, DNA)
      }

      DNA = v
    } else {
      introns = append(introns, v)
    }
  }

  return DNA, introns
}


func remove_introns(DNA string, introns []string) (string) {

  for _, intron := range introns {
    DNA = strings.Replace(DNA, intron, "", -1)
  }

  return DNA
}


func DNA2RNA(DNA string) (string) {

  RNA := strings.Replace(DNA, "T", "U", -1)
  return RNA
}


func translate(sequence string, table map[string]string) (string, bool) {

  ul := 3 // unit length
  protein_sequence := ""
  iter_num := (len(sequence) - len(sequence) % ul) / ul

  for i := 0; i < iter_num; i++ {
    protein := table[sequence[ul * i:ul * (i + 1)]]
    if protein == "" {
      return "", false
    } else if protein == "Stop" {
      break
    } else {
      protein_sequence += protein
    }
  }

  return protein_sequence, true
}


func main() {

  text, is_text_success := load_text("rosalind_splc.txt")
  RNA_codon_table, is_table_success := load_table("RNA_codon_table.txt")

  if is_text_success && is_table_success {
    DNAs := load_samples(text)
    DNA, introns := select_DNA_introns(DNAs)
    DNA_extrons := remove_introns(DNA, introns)

    RNA := DNA2RNA(DNA_extrons)
    protein, is_success := translate(RNA, RNA_codon_table)
    if is_success {
      fmt.Println(protein)
    }
  }
}