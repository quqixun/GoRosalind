package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func load_DNA(file_path string) (string, bool) {

  DNA := ""

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return DNA, false
  }

  scanner := bufio.NewScanner(strings.NewReader(string(b)))
  for scanner.Scan() {
    if strings.Index(scanner.Text(), ">") == -1 {
      DNA += scanner.Text()
    }
  }

  return DNA, true
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


func reverse_palindrome(DNA, DNA_rc string) ([][]int) {

  rp := [][]int{}

  DNA_len := len(DNA)
  for l := 4; l <= 12; l++ {
    for i := 0; i < DNA_len - l + 1; i++ {
      if DNA[i:i + l] == DNA_rc[DNA_len - i - l:DNA_len - i] {
        rp = append(rp, []int{i + 1, l})
        // fmt.Println(i + 1, l)
      }
    }
  }

  return rp
}


func main() {

  DNA, is_load_success := load_DNA("rosalind_revp.txt")

  if is_load_success {
    DNA_rc := reverse_complement(DNA)
    rp := reverse_palindrome(DNA, DNA_rc)

    for _, r := range rp {
      fmt.Println(strconv.Itoa(r[0]) + " " + strconv.Itoa(r[1]))
    }
  }
}