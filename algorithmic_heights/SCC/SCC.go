package main

import (
  "fmt"
  // "math"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, [][]int) {

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

  n_strs := strings.Split(content[0], " ")
  n, _ := strconv.Atoi(n_strs[0])

  graph := [][]int{}
  for _, c := range content[1:len(content)] {
    edge := []int{}
    num_strs := strings.Split(c, " ")
    for _, num_str := range num_strs {
      num, _ := strconv.Atoi(num_str)
      edge = append(edge, num)
    }
    graph = append(graph, edge)
  }

  return n, graph
}


func main() {

  n, graph := LoadData("rosalind_scc.txt")
  fmt.Println(n, graph)

}