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


func BreadthFirstSearch(n_vertices int, edges [][]int) ([]int) {

  res := []int{}
  for i := 0; i < n_vertices; i++ {
    res = append(res, -1)
  }
  res[0] = 0

  distance := 0
  nodes := []int{1}
  visited := []int{1}
  for {
    if len(nodes) == 0 {
      break
    } else {
      for _, n := range nodes {
        res[n - 1] = distance
      }

      new_nodes := []int{}
      for _, n := range nodes {
        for _, edge := range edges {
          if edge[0] == n {
            if !IsInSlice(visited, edge[1]) {
              visited = append(visited, edge[1])
              new_nodes = append(new_nodes, edge[1])            
            }
          }
        }
      }

      nodes = new_nodes
      distance += 1
    }
  }

  return res
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

  n_vertices, _, edges := LoadData("rosalind_bfs.txt")
  res := BreadthFirstSearch(n_vertices, edges)
  fmt.Println(IntSliceToString(res))
}