package main

import (
  "fmt"
  "math"
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


func LoadData(file_path string) (int, int, []int, []int) {

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
  m, _ := strconv.Atoi(content[1])
  A := String2IntSlice(content[2])
  K := String2IntSlice(content[3])

  return n, m, A, K
}


func BinarySearch(A []int, k int) (int) {

  index := -1
  low, high := 0, len(A)

  for {
    if low >= high {
      break
    }

    mid := int(math.Floor(float64(low + high) / 2.0))
    mid_value := A[mid]
    if mid_value < k {
      low = mid + 1
    } else if mid_value > k {
      high = mid
    } else {
      index = mid + 1
      break
    }
  }

  return index
}


func IntSliceToString(ints []int) (string) {

  int_strs := []string{}
  for _, p := range ints {
    int_strs = append(int_strs, strconv.Itoa(p))
  }

  str := strings.Join(int_strs, " ")
  return str
}


func main() {
  
  _, _, A, K := LoadData("rosalind_bins.txt")
  
  indices := []int{}
  for _, k := range K {
    indices = append(indices, BinarySearch(A, k))
  }
  fmt.Println(IntSliceToString(indices))
}