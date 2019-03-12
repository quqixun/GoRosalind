package main

import (
  "fmt"
  "strconv"
  "strings"
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


func Fibonacci(n int) (int) {
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  }

  return Fibonacci(n - 1) + Fibonacci(n - 2)
}


func main() {
  
  n := LoadData("rosalind_fibo.txt")
  fibo := Fibonacci(n)
  fmt.Println(fibo)
}