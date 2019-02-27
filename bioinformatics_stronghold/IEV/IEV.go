package main

import (
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
)


func read_numbers(file_path string) ([]int, bool) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return nil, false
  }

  text := string(b)
  words := strings.Fields(text)

  numbers := []int{}
  for _, w := range words{
    n, err_n := strconv.Atoi(w)
    if err_n != nil {
      fmt.Println(err_n)
      return nil, false
    }

    numbers = append(numbers, n)
  }

  return numbers, true
}


func expected_offspring(numbers []int) (float64) {

  n_offspring := 2.0
  prop := []float64{1.0, 1.0, 1.0, 0.75, 0.5, 0}

  E := 0.0
  for i := 0; i < len(numbers); i++ {
    E += float64(numbers[i]) * n_offspring * prop[i]
  }

  return E
}


func main() {

  numbers, is_read_success := read_numbers("rosalind_iev.txt")

  if is_read_success {
    E := expected_offspring(numbers)
    fmt.Println(E)
  }
}