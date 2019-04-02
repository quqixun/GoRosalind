package main

import (
  "fmt"
  "math"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, string, []float64) {

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
  s := content[1]
  nums := strings.Fields(content[2])

  A := []float64{}
  for _, n := range nums {
    f, _ := strconv.ParseFloat(n, 64)
    A = append(A, f)
  }

  return n, s, A
}


func Expection(n int, s string, gc_content float64) (float64) {

  gc_prob := gc_content / 2.0
  at_prob := (1 - gc_content) / 2.0

  nA := float64(strings.Count(s, "A"))
  nC := float64(strings.Count(s, "C"))
  nG := float64(strings.Count(s, "G"))
  nT := float64(strings.Count(s, "T"))

  l := float64(n - len(s) + 1)
  exp := math.Pow(at_prob, nA) *
         math.Pow(at_prob, nT) *
         math.Pow(gc_prob, nC) *
         math.Pow(gc_prob, nG) * l

  return exp
}


func FloatsToString(res []float64) (string) {

  res_strs := []string{}
  for _, r := range res {
    res_strs = append(res_strs, fmt.Sprintf("%.3f", r))
  }

  res_str := strings.Join(res_strs, " ")
  return res_str
}


func main() {
  
  n, s, A := LoadData("rosalind_eval.txt")

  res := []float64{}
  for _, gc_content := range A {
    res = append(res, Expection(n, s, gc_content))
  }
  fmt.Println(FloatsToString(res))
}