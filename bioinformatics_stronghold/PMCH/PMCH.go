/*
  Perfect Matchings of Bipartite Graph
  Explanation:
  (Chinese) https://www.renfei.org/blog/bipartite-matching.html
*/

package main

import (
  "fmt"
  "bufio"
  "strings"
  "math/big"
  "io/ioutil"
)


func load_RNA(file_path string) (string, bool) {

  RNA := ""

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return RNA, false
  }

  scanner := bufio.NewScanner(strings.NewReader(string(b)))
  for scanner.Scan() {
    if strings.Index(scanner.Text(), ">") == -1 {
      RNA += scanner.Text()
    }
  }

  return RNA, true
}


func factorial(n int64) (r *big.Int) {

  if n > 0 {
    big_n := big.NewInt(n)
    r = big_n.Mul(big_n, factorial(n - 1))
    return r
  }

  return big.NewInt(1)
}


func perfect_matchings(RNA string) (*big.Int) {

  n_A := int64(strings.Count(RNA, "A"))
  n_C := int64(strings.Count(RNA, "C"))

  fac_A := factorial(n_A)
  fac_C := factorial(n_C)
  n_pm := fac_A.Mul(fac_A, fac_C)
  return n_pm
}


func main() {

  RNA, is_load_success := load_RNA("rosalind_pmch.txt")

  if is_load_success {
    n_pm := perfect_matchings(RNA)
    fmt.Printf(n_pm.String())
  }
}