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


func BuildHeap(A []int, i int) ([]int) {

  if i == 0 {
    return A
  }

  parent, child := (i - 1) / 2, i
  if A[parent] >= A[child] {
    return A
  } else {
    A[parent], A[child] = A[child], A[parent]
    return BuildHeap(A, parent)
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
  
  _, A := LoadData("rosalind_hea.txt")
  
  pA := []int{}
  for i, a := range A {
    pA = append(pA, a)
    pA = BuildHeap(pA, i)
  }

  Write2File("hea_output.txt", []string{IntSliceToString(pA)})
}