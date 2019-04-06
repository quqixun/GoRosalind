package main

import (
  "fmt"
  "bufio"
  "strconv"
  "strings"
  "io/ioutil"
)


func LoadData(file_path string) (int, [][][]int) {

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

  graphs, graph := [][][]int{}, [][]int{}
  for _, c := range content[1:len(content)] {
    if c == "" {
      if len(graph) > 0 {
        graphs = append(graphs, graph)
        graph = [][]int{}
      }
    } else {
      num_strs := strings.Split(c, " ")
      nums := []int{}
      for _, num_str := range num_strs {
        num, _ := strconv.Atoi(num_str)
        nums = append(nums, num)
      }
      graph = append(graph, nums)
    }
  }
  graphs = append(graphs, graph)

  return n, graphs
}


func AdjacentList(n_vertices int, edges [][]int) ([][]int) {

  adjacent := [][]int{}
  for i := 0; i < n_vertices; i++ {
    adjacent = append(adjacent, []int{})
  }

  for _, e := range edges {
    adjacent[e[0] - 1] = append(adjacent[e[0] - 1], e[1] - 1)
    adjacent[e[1] - 1] = append(adjacent[e[1] - 1], e[0] - 1)
  }

  return adjacent
}


func IsBipartiteGraph(adjacent [][]int, v int, visited []bool, color []int) (bool) {

  for _, u := range adjacent[v] {
    if !visited[u] {
      visited[u] = true
      color[u] = 1 - color[v]

      if !IsBipartiteGraph(adjacent, u, visited, color){
        return false
      }
    } else if color[u] == color[v] {
      return false
    }
  }

  return true
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
  
  _, graphs := LoadData("rosalind_bip.txt")
  
  res := []int{}
  for _, graph := range graphs {

    n_vertices := graph[0][0]
    edges := graph[1:len(graph)]

    color := []int{}
    for i := 0; i < n_vertices; i++ {
      color = append(color, -1)
    }
    visited := make([]bool, n_vertices)

    color[edges[0][0]] = 0
    visited[edges[0][0]] = true
    adjacent := AdjacentList(n_vertices, edges)
    resb := IsBipartiteGraph(adjacent, 0, visited, color)

    var resi int
    if resi = -1; resb {
      resi = 1
    }
    res = append(res, resi)
  }
  fmt.Println(IntSliceToString(res))
}