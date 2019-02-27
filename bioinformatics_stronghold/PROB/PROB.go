package main

import (
  "fmt"
  "math"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func load_data(file_path string) (string, []float64, bool) {

  DNA := ""
  A := []float64{}

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return DNA, A, false
  }

  lines := []string{}
  scanner := bufio.NewScanner(strings.NewReader(string(b)))
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  DNA = lines[0]
  A_words := strings.Fields(lines[1])
  for _, word := range A_words {
    f, _ := strconv.ParseFloat(word, 64)
    A = append(A, f)
  }

  return DNA, A, true
}


func common_logarithm(DNA string, A []float64) ([]float64) {

  cl := []float64{}

  for _, x := range A {
    cl_x := 0.0
    for _, c := range DNA {
      c_str := string(c)
      switch c_str {
        case "G": cl_x += math.Log10(x * 0.5)
        case "C": cl_x += math.Log10(x * 0.5)
        case "A": cl_x += math.Log10((1 - x) * 0.5)
        case "T": cl_x += math.Log10((1 - x) * 0.5)
      }
    }
    cl = append(cl, cl_x)
  }

  return cl
}


func main() {

  DNA, A, is_load_success := load_data("rosalind_prob.txt")

  if is_load_success {
    cl := common_logarithm(DNA, A)
    
    cl_strs := []string{}
    for _, p := range cl {
      cl_strs = append(cl_strs, fmt.Sprintf("%.3f", p))
    }
    cl_str := strings.Join(cl_strs, " ")
    fmt.Println(cl_str)
  }
}