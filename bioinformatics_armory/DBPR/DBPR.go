package main

import (
  "io"
  "os"
  "fmt"
  "bufio"
  "strings"
  "net/http"
  "io/ioutil"
)


func load_protein(file_path string) (string, bool) {

  protein := ""

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return protein, false
  }

  protein = string(b)
  return protein, true
}


func download_file(protein, save_path string) (error) {

  url := "http://www.uniprot.org/uniprot/" + protein + ".txt"

  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  out, err := os.Create(save_path)
  if err != nil {
    return err
  }
  defer out.Close()

  _, err = io.Copy(out, resp.Body)
  return err
}


func load_text(file_path string) ([]string) {

  content := []string{}

  b, err := ioutil.ReadFile(file_path)
  if err != nil {
    fmt.Println(err)
    return content
  }

  text := string(b)
  scanner := bufio.NewScanner(strings.NewReader(text))
  for scanner.Scan() {
    content = append(content, scanner.Text())
  }

  return content
}


func find_processes(content []string) ([]string) {

  processes := []string{}

  for _, c := range content {
    if c[:2] != "DR" {
      continue
    } else {
      P_index := strings.Index(c, "P:")
      if P_index > -1 {
        S_index := -1
        for i := P_index; i < len(c); i++ {
          if string(c[i]) == ";" {
            S_index = i
          }
        }
        processes = append(processes, c[P_index + 2:S_index])
      }
    }
  }

  return processes
}


func main() {
  
  protein, is_load_success := load_protein("rosalind_dbpr.txt")

  if is_load_success {
    save_path := protein + ".txt"
    err := download_file(protein, save_path)
    if err != nil {
      panic(err)
    }

    content := load_text(save_path)
    processes := find_processes(content)
    for _, p := range processes {
      fmt.Println(p)
    }
  }
}