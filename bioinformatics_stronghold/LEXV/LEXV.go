package main

import (
  "os"
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) ([]string, int) {

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

  A := strings.Fields(content[0])
  n, _ := strconv.Atoi(content[1])

  return A, n
}


func SortedKMers(A []string, n int,
                 item string, pres *[]string) () {

  if n > 0 {
    for _, c := range A {
      *pres = append(*pres, item + c)
      SortedKMers(A, n - 1, item + c, pres)
    }
  }

  return
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
  
  A, n := LoadData("rosalind_lexv.txt")
  res := []string{}
  SortedKMers(A, n, "", &res)
  Write2File("lexv_output.txt", res)
}