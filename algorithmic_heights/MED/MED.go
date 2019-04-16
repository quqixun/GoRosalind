package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, []int, int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := []string{}
  scanner := bufio.NewScanner(strings.NewReader(text))

  buf := make([]byte, 0, 64*1024)
  scanner.Buffer(buf, 1024*1024)

  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  n, _ := strconv.Atoi(content[0])
  num_strs := strings.Split(content[1], " ")
  k, _ := strconv.Atoi(content[2])

  A := []int{}
  for _, num_str := range num_strs {
    num, _ := strconv.Atoi(num_str)
    A = append(A, num)
  }

  return n, A, k
}


func SelectionSort(A []int) {

  for i := 0; i < len(A); i++ {
    min_idx := i
    for j := i + 1; j < len(A); j++ {
      if A[min_idx] > A[j] {
        min_idx = j
      }
    }

    A[i], A[min_idx] = A[min_idx], A[i]
  }
}


func main() {

  _, A, k := LoadData("rosalind_med.txt")

  SelectionSort(A)
  fmt.Println(A[k - 1])
}