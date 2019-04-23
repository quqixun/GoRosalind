package main

import (
  "os"
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
  }

  return adjacent
}


func IsInSlice(slice []int, i int) (bool) {

  for _, s := range slice {
    if s == i {
      return true
    }
  }

  return false
}


func TopologicalSortUtil(adj [][]int, v int, visited []bool, stack []int) ([]int) {
  visited[v] = true
  for _, i := range adj[v] {
    if visited[i] == false {
      stack = TopologicalSortUtil(adj, i, visited, stack)
    }
  }
  stack = append([]int{v + 1}, stack...)
  return stack
}


func HamiltonianCycle(n_vertices int, adjacent [][]int, visited []bool, stack []int) (int, []int) {

  for i := 0; i < n_vertices; i++ {
    if visited[i] == false {
      stack = TopologicalSortUtil(adjacent, i, visited, stack)
    }
  }

  for i := 0; i < len(stack) - 1; i++ {
    idx1, idx2 := stack[i] - 1, stack[i + 1] - 1
    if !IsInSlice(adjacent[idx1], idx2) {
      return -1, []int{}
    }
  }

  return 1, stack
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

  _, graphs := LoadData("rosalind_hdag.txt")

  result_content := []string{}
  for _, graph := range graphs {
    n_vertices := graph[0][0]
    edges := graph[1:len(graph)]

    adjacent := AdjacentList(n_vertices, edges)
    visited := make([]bool, n_vertices)
    stack := make([]int, 0)

    res, path := HamiltonianCycle(n_vertices, adjacent, visited, stack)
    result := append([]int{res}, path...)
    result_str := IntSliceToString(result)
    result_content = append(result_content, result_str)
  }
  Write2File("hdag_output.txt", result_content)
}