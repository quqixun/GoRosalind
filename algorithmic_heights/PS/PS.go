package main

import (
  "os"
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


func Heapify(A []int, n, i int) {

  largest := i
  left := 2 * i + 1
  right := 2 * i + 2

  if left < n && A[i] < A[left] {
    largest = left
  }

  if right < n && A[largest] < A[right] {
    largest = right
  }

  if largest != i {
    A[i], A[largest] = A[largest], A[i]
    Heapify(A, n, largest)
  }
}


func HeapSort(A []int) {

  n := len(A)

  for i := n; i >= 0; i-- {
    Heapify(A, n, i)
  }

  for i := n - 1; i >= 1; i-- {
    A[i], A[0] = A[0], A[i]
    Heapify(A, i, 0)
  }
}


func IntSliceToString(ints []int) (string) {

  int_strs := []string{}
  for _, p := range ints {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func Write2File(file_path string, content []string) {
  f, err := os.Create(file_path)
  if err != nil {
      panic(err)
  }
  defer f.Close()

  for _, c := range content {
     fmt.Fprintln(f, c)
  }
}


func main() {
  
  _, A, k := LoadData("rosalind_ps.txt")
  
  HeapSort(A)
  Write2File("ps_output.txt", []string{IntSliceToString(A[0:k])})
}