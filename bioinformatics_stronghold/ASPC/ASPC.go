package main

import (
  "fmt"
  "strconv"
  "strings"
  "math/big"
  "io/ioutil"
)


func LoadData(file_path string) (int, int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := strings.TrimSuffix(string(b), "\r\n")
  num_strs := strings.Split(text, " ")

  n, _ := strconv.Atoi(num_strs[0])
  m, _ := strconv.Atoi(num_strs[1])
  return n, m
}


func Factorial(n int) (r *big.Int) {

  if n > 0 {
    big_n := big.NewInt(int64(n))
    r = big_n.Mul(big_n, Factorial(n - 1))
    return r
  }

  return big.NewInt(int64(1))
}


func CountSubset(n, m int) (*big.Int) {

  count := big.NewInt(0)
  for k := m; k <= n; k++ {
    count_k := Factorial(n)
    count_k = count_k.Div(count_k, Factorial(k))
    count_k = count_k.Div(count_k, Factorial(n - k))
    count = count.Add(count, count_k)
  }

  return count
}


func main() {
  
  n, m := LoadData("rosalind_aspc.txt")
  count := CountSubset(n, m)
  count = count.Mod(count, big.NewInt(1000000))
  fmt.Println(count.String())
}