package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, []int) {

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

  A := []int{}
  for _, num_str := range num_strs {
    num, _ := strconv.Atoi(num_str)
    A = append(A, num)
  }

  return n, A
}


func Partition(A []int, low, high int) (int) {

  i := low - 1
  pivot := A[high]

  for j := low; j < high; j++ {
    if A[j] < pivot {
      i += 1
      A[i], A[j] = A[j], A[i]
    }
  }

  A[i + 1], A[high] = A[high], A[i + 1]
  return i + 1
}


func QuickSort(A []int, low, high int) {

  if low < high {
    pi := Partition(A, low, high)
    QuickSort(A, low, pi - 1)
    QuickSort(A, pi + 1, high)
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

  n, A := LoadData("rosalind_qs.txt")

  QuickSort(A, 0, n - 1)
  Write2File("qs_output.txt", []string{IntSliceToString(A)})
}