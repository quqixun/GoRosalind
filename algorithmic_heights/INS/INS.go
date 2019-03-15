package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func String2IntSlice(str string) ([]int) {

  ints := []int{}
  int_strs := strings.Fields(str)
  for _, is := range int_strs {
    i, _ := strconv.Atoi(is)
    ints = append(ints, i)
  }

  return ints
}


func LoadData(file_path string) (int, []int) {

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
  A := String2IntSlice(content[1])

  return n, A
}


func InsertionSort(A []int) (int) {

  n_swap := 0

  for i := 1; i <= len(A); i++ {
    k := i
    for {
      if k > 1 && A[k - 1] < A[k - 2] {
        temp := A[k - 1]
        A[k - 1] = A[k - 2]
        A[k - 2] = temp
        n_swap += 1
        k -= 1
      } else {
        break
      }
    }
  }

  return n_swap
}


func main() {
  
  _, A := LoadData("rosalind_ins.txt")

  n_swap := InsertionSort(A)
  fmt.Println(n_swap)
}