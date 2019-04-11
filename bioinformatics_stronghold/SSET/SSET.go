package main

import (
  "fmt"
  "math"
  "strconv"
  "strings"
  "math/big"
  "io/ioutil"
)


func LoadData(file_path string) (int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := strings.TrimSuffix(string(b), "\r\n")
  n, _ := strconv.Atoi(text)
  return n
}


func Factorial(n int) (r *big.Int) {

  if n > 0 {
    big_n := big.NewInt(int64(n))
    r = big_n.Mul(big_n, Factorial(n - 1))
    return r
  }

  return big.NewInt(int64(1))
}


func CountSubset(n int) (*big.Int) {

  m := -1
  mod := math.Mod(float64(n), 2.0)
  if mod == 0.0 {
    m = n / 2
  } else {
    m = (n - 1) / 2
  }

  count := big.NewInt(2)
  for i := 1; i <= m; i++ {
    count_i := Factorial(n)
    count_i = count_i.Div(count_i, Factorial(i))
    count_i = count_i.Div(count_i, Factorial(n - i))

    dcount_i := new(big.Int).Set(count_i)
    dcount_i = dcount_i.Mul(dcount_i, big.NewInt(int64(2)))

    if i != m {
      count = count.Add(count, dcount_i)
    } else {
      if mod == 0.0 {
        count = count.Add(count, count_i)
      } else {
        count = count.Add(count, dcount_i)
      }
    }
  }

  return count
}


func main() {
  
  n := LoadData("rosalind_sset.txt")
  count := CountSubset(n)
  count = count.Mod(count, big.NewInt(1000000))
  fmt.Println(count.String())
}