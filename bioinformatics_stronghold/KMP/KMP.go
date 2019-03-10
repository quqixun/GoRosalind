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

  seq := ""
  text := string(b)
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    line_text := scanner.Text()
    if strings.Index(line_text, ">") == -1 {
      seq += line_text
    }
  }

  return seq
}


func  KnuthMorrisPratt(seq string) ([]int) {

  P := make([]int, len(seq))

  for k := 1; k <= len(seq); k++ {
    if k == 1{
      P[k - 1] = 0
    } else {
      sub_seq := seq[0:k]
      for j := k - 1; j >= 1; j-- {
        prefix := sub_seq[0:k - j]
        suffix := sub_seq[j:k]
        if prefix == suffix {
          if len(suffix) > P[k - 1]{
            P[k - 1] = len(suffix)
          }
        }
      } 
    }
  }

  return P
}


func IntSliceToString(ints []int) (string) {

  int_strs := []string{}
  for _, p := range ints {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func main() {
  
  seq := LoadData("rosalind_kmp.txt")
  P := KnuthMorrisPratt(seq)
  fmt.Println(IntSliceToString(P))

}