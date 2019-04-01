package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
)


func LoadData(file_path string) (int, []int) {

  f, err := os.Open(file_path)
  if err != nil {
    panic(err)
  }

  content := [][]int{}
  scanner := bufio.NewScanner(f)
  buf := make([]byte, 0, 64*1024)
  scanner.Buffer(buf, 1024*1024)
  for scanner.Scan() {
    int_strs := strings.Fields(scanner.Text())
    ints := []int{}
    for _, str := range int_strs {
      i, _ := strconv.Atoi(str)
      ints = append(ints, i)
    }
    content = append(content, ints)
  }

  n := content[0][0]
  A := content[1]
  return n, A
}


func CountingInversions(A []int, n int) (int) {

  n_inv := 0
  for i := 0; i < n - 1; i++ {
    for j := i + 1; j < n; j++ {
      if A[i] > A[j] {
        n_inv += 1
      }
    }
  }

  return n_inv
}


func main() {
  n, A := LoadData("rosalind_inv.txt")
  n_inv := CountingInversions(A, n)
  fmt.Println(n_inv)
}