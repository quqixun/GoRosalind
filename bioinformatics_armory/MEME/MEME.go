package main

import (
  "os"
  "fmt"
  "strconv"
  "os/exec"
)


func MeMe(file_path string, min_lenfth int) (string) {

  cmd := exec.Command("meme", file_path,
                      "-text", "-nostatus", "-protein",
                      "-minw", strconv.Itoa(min_lenfth))

  output, err := cmd.CombinedOutput()
  if err != nil {
    panic(err)
  }
  
  return string(output)
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

    output := MeMe("rosalind_meme.txt", 20)
    Write2File("meme_output.txt", []string{output})
}