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

  text := string(b)
  seqs := []string{}
  scanner := bufio.NewScanner(strings.NewReader(text))

  seq := ""
  for scanner.Scan() {
    line_text := scanner.Text()
    if strings.Index(line_text, ">") > -1 {
      if seq != "" {
        seqs = append(seqs, seq)
        seq = ""
      }
    } else {
      seq += line_text
    }
  }

  seqs = append(seqs, seq)
  return seqs
}


func ReverseComplement(seq string) (string) {

  rc_seq := ""
  for _, c := range seq {
    sc := string(c)
    switch sc {
      case "A": rc_seq = "T" + rc_seq
      case "T": rc_seq = "A" + rc_seq
      case "C": rc_seq = "G" + rc_seq
      case "G": rc_seq = "C" + rc_seq
    }
  }

  return rc_seq
}


func Count(slice []string, target string) (int) {

  num := 0
  for _, s := range slice {
    if s == target {
      num += 1
    }
  }

  return num
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


func FindCorrectAndIncorrect(seqs []string) ([]string, []string) {

  num := 0
  correct := []string{}
  incorrect := []string{}

  for _, seq := range seqs {
    seq_rc := ReverseComplement(seq)
    if seq == seq_rc {
      num = Count(seqs, seq)
    } else {
      num = Count(seqs, seq) + Count(seqs, seq_rc)
    }

    if num >= 2 &&
       Count(correct, seq) == 0 &&
       Count(correct, seq_rc) == 0 {
      correct = append(correct, seq)
    } else {
      incorrect = append(incorrect, seq)
    }
  }

  return correct, incorrect
}


func ErrorCorrection(seqs []string) ([][]string) {

  correct, incorrect := FindCorrectAndIncorrect(seqs)

  corrs := [][]string{}
  for _, incorr := range incorrect {
    for _, corr := range correct {
      corr_rc := ReverseComplement(corr)
      if HammingDistance(incorr, corr) == 1 {
        corrs = append(corrs, []string{incorr, corr})
        break
      } else if HammingDistance(incorr, corr_rc) == 1 {
        corrs = append(corrs, []string{incorr, corr_rc})
        break
      } else { continue }
    }
  }

  return corrs
}


func main() {
  
  seqs := LoadData("rosalind_corr.txt")
  corrs := ErrorCorrection(seqs)
  
  for _, corr := range corrs {
    fmt.Printf("%s->%s\r\n", corr[0], corr[1])
  }
}