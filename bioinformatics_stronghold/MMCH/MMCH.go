package main

import (
  "fmt"
  "bufio"
  "strings"
  "math/big"
  "io/ioutil"
)


func LoadData(file_path string) (string) {

  RNA := ""

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  scanner := bufio.NewScanner(strings.NewReader(string(b)))
  for scanner.Scan() {
    if strings.Index(scanner.Text(), ">") == -1 {
      RNA += scanner.Text()
    }
  }

  return RNA
}


func Factorial(n int64) (r *big.Int) {

  if n > 0 {
    big_n := big.NewInt(n)
    r = big_n.Mul(big_n, Factorial(n - 1))
    return r
  }

  return big.NewInt(1)
}


func PairMatchings(RNA string, pairs string) (*big.Int) {

  n1, n2 := int64(0), int64(0)
  if pairs == "AU" {
    n1 = int64(strings.Count(RNA, "A"))
    n2 = int64(strings.Count(RNA, "U"))
  } else {
    n1 = int64(strings.Count(RNA, "G"))
    n2 = int64(strings.Count(RNA, "C"))
  }

  max, min := n1, n2
  if n2 > n1 {
    max, min = n2, n1
  }
  diff := max - min

  n := Factorial(max)
  d := Factorial(diff)
  m := n.Div(n, d)
  return m
}


func MaximumMatchings(RNA string) (*big.Int) {

  aum := PairMatchings(RNA, "AU")
  gcm := PairMatchings(RNA, "GC")

  n_mm := aum.Mul(aum, gcm)
  return n_mm
}


func main() {

  RNA := LoadData("rosalind_mmch.txt")

  n_mm := MaximumMatchings(RNA)
  fmt.Printf(n_mm.String())

}