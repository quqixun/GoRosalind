package main

import (
  "fmt"
  "math"
  "strconv"
  "strings"
  "io/ioutil"
)


func load_kN(file_path string) (uint64, uint64, bool) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return 0, 0, false
  }

  text := string(b)
  numbers := strings.Fields(text)

  k, _ := strconv.Atoi(numbers[0])
  N, _ := strconv.Atoi(numbers[1])
  return uint64(k), uint64(N), true
}


func factorial(n uint64) (r float64) {

  if n > 0 {
    r = float64(n) * factorial(n - 1)
    return r
  }

  return 1
}


func combinations(pop, n uint64) (float64) {

  pop_fac := factorial(pop)
  n_fac := factorial(n)
  pop_n_fac := factorial(pop -n)

  n_com := pop_fac / n_fac / pop_n_fac
  return n_com
}


func sum(slice []float64) (float64) {

  r := 0.0
  for  _, s := range slice {
    r += s
  }

  return r
}


func independent_alleles(k, N uint64) (float64) {

  pop := uint64(math.Pow(2.0, float64(k)))
  P := []float64{1.0, 0.0, 0.0}
  P_temp := []float64{0.0, 0.0, 0.0}

  var i uint64
  for i = 1; i <= k; i++ {
    P_temp[0] = 0.5 * sum(P)
    P_temp[1] = 0.25 * P[0] + 0.5 * P[1]
    P_temp[2] = 0.25 * P[0] + 0.5 * P[2]
    copy(P, P_temp)
  }

  prob := 0.0
  P_AaBb := math.Pow(P[0], 2.0)

  var n uint64
  for n = N; n <= pop; n++ {
    prob += combinations(pop, n) *
            math.Pow(P_AaBb, float64(n)) *
            math.Pow(1.0 - P_AaBb, float64(pop - n))
  }

  return prob
}


func main() {

  k, N, is_load_success := load_kN("rosalind_lia.txt")

  if is_load_success {
    prob := independent_alleles(k, N)
    fmt.Println((prob))
  }
}