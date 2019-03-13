package main

import (
  "os"
  "fmt"
  "strconv"
  "strings"
  "io/ioutil"
)


func String2IntSlice(str string) ([]int) {

  str = str[1:len(str) - 1]
  int_strs := strings.Split(str, ", ")

  ints := []int{}
  for _, int_str := range int_strs {
    int_n, _ := strconv.Atoi(int_str)
    ints = append(ints, int_n)
  }

  return ints
}


func LoadData(file_path string) (int, []int, []int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := strings.Split(text, "\r\n")

  n, _ := strconv.Atoi(content[0])
  A := String2IntSlice(content[1])
  B := String2IntSlice(content[2])

  return n, A, B
}


func Range(n int) ([]int) {

  ints := []int{}
  for i := 1; i <= n; i++ {
    ints = append(ints, i)
  }

  return ints
}


func IntInSlice(num int, slice []int) (bool) {

  for _, i := range slice {
    if num == i {
      return true
    }
  }

  return false
}


func SetUnion(A, B []int) ([]int) {

  union := make([]int, len(A))
  copy(union, A)

  for _, b := range B {
    if !IntInSlice(b, union) {
      union = append(union, b)
    }
  }

  return union
}


func SetIntersection(A, B []int) ([]int) {

  intersection := []int{}

  for _, a := range A {
    for _, b := range B {
      if a == b {
        intersection = append(intersection, a)
      }
    }
  }

  return intersection
}


func SetDifference(A, B []int) ([]int) {

  difference := []int{}

  for _, a := range A {
    if !IntInSlice(a, B) {
      difference = append(difference, a)
    }
  }

  return difference
}


func SetComplement(A, U []int) ([]int) {

  complement := SetDifference(U, A)

  return complement
}


func IntSliceToSetString(ints []int) (string) {

  int_strs := []string{}
  for _, p := range ints {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, ", ")
  str = "{" + str + "}"
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
  
  n, A, B := LoadData("rosalind_seto.txt")
  U := Range(n)

  union := SetUnion(A, B)
  intersection := SetIntersection(A, B)
  AB_difference := SetDifference(A, B)
  BA_difference := SetDifference(B, A)
  A_complement := SetComplement(A, U)
  B_complement := SetComplement(B, U)

  result := []string{
    IntSliceToSetString(union),
    IntSliceToSetString(intersection),
    IntSliceToSetString(AB_difference),
    IntSliceToSetString(BA_difference),
    IntSliceToSetString(A_complement),
    IntSliceToSetString(B_complement),
  }
  Write2File("seto_output.txt", result)
}