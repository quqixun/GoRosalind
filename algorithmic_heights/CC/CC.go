package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, int, [][]int) {

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    panic(err)
  }

  text := string(b)
  content := [][]int{}
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    int_strs := strings.Fields(scanner.Text())
    ints := []int{}
    for _, str := range int_strs {
      i, _ := strconv.Atoi(str)
      ints = append(ints, i)
    }
    content = append(content, ints)
  }

  n_vertices := content[0][0]
  n_edges := content[0][1]
  edges := content[1:len(content)]
  return n_vertices, n_edges, edges
}


func IsInSlice(ints []int, t int) (bool) {

  for _, i := range ints {
    if i == t {
      return true
    }
  }

  return false
}


func DepthFirstSearch(adjacent_edges [][]int, visited map[int]bool, i int) {

  visited[i] = true

  for _, adjacent := range adjacent_edges[i] {
    if !visited[adjacent - 1] {
      DepthFirstSearch(adjacent_edges, visited, adjacent - 1)
    }
  }

  return
}


func CountConnectedComponents(n_vertices, n_edges int, edges [][]int) (int) {

  adjacent_edges := [][]int{}
  for i := 0; i < n_vertices; i++ {
    adjacent_edges = append(adjacent_edges, []int{})
  }

  for _, e := range edges {
    adjacent_edges[e[0] - 1] = append(adjacent_edges[e[0] - 1], e[1])
    adjacent_edges[e[1] - 1] = append(adjacent_edges[e[1] - 1], e[0])
  }

  n_components := 0
  visited := map[int]bool{}
  for i := 0; i < n_vertices; i++ {
    visited[i] = false
  }

  for i := 0; i < n_vertices; i++ {
    if !visited[i] {
      n_components += 1
      DepthFirstSearch(adjacent_edges, visited, i)
    }
  }

  return n_components
}


func main() {

  n_vertices, n_edges, edges := LoadData("rosalind_cc.txt")
  cc := CountConnectedComponents(n_vertices, n_edges, edges)
  fmt.Println(cc)
}