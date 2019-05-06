package main

import (
  "fmt"
  "math"
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
      continue
    }
    num_strs := strings.Split(c, " ")
    nums := []int{}
    for _, num_str := range num_strs {
      num, _ := strconv.Atoi(num_str)
      nums = append(nums, num)
    }

    if len(nums) == 2 {
      if len(graph) > 0 {
        graphs = append(graphs, graph)
      }
      graph = [][]int{nums}
    } else {
      graph = append(graph, nums)
    }
  }
  graphs = append(graphs, graph)

  return n, graphs
}


func IsVisited(i int, visited []int) (bool) {

  for _, v := range visited {
    if i == v {
      return true
    }
  }

  return false
}


func NegatuceWeightCycle(n int, edges [][]int) (bool) {

  sources := []int{}
  visited := []int{}

  for i := 0; i < n; i++ {
    sources = append(sources, i)
  }

  for _, source := range sources {

    if IsVisited(source, visited) {
      continue
    }

    dist := []float64{}
    for i := 0; i < n; i++ {
      dist = append(dist, math.Inf(1))
    }
    dist[source] = 0
    changes := true

    for i := 0; i < len(edges); i++ {
      if !changes {
        break
      }
      changes = false

      extra := false
      if i == len(edges) - 1 {
        extra = true
      }

      for _, edge := range edges {
        u, v, w := edge[0], edge[1], float64(edge[2])
        if dist[u - 1] != math.Inf(1) &&
           (dist[v - 1] == math.Inf(1) || dist[u - 1] + w < dist[v - 1]) {
          dist[v - 1] = dist[u - 1] + w
          changes = true
          if extra {
            return true
          }
        }
      }
    }

    for i, d := range dist {
      if d != math.Inf(1) {
        visited = append(visited, i)
      }
    }
  }

  return false
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

  _, graphs := LoadData("rosalind_nwc.txt")
  
  result := []int{}
  for _, graph := range graphs {
    n := graph[0][0]
    edges := graph[1:len(graph)]

    if NegatuceWeightCycle(n, edges) {
      result = append(result, 1)
    } else {
      result = append(result, -1)
    }
  }

  fmt.Println(IntSliceToString(result))
}