package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
  "math"
)

var store = map[string][]string{}

func polyHash(s string, bsize int) string {
  hash := 0
  for i := len(s) - 1; i >= 0; i-- {
    hash = ((263 * hash) + int(s[i])) % (int(math.Pow10(9)) + 7)
  }
  return strconv.Itoa(hash % bsize)
}

func containsString(s string) bool {
  _, contains := store[s]
  if contains {
    return true
  }
  return false
}

func insideSlice(slice []string, target string) bool {
  for _, v := range slice {
    if v == target {
      return true
    }
  }
  return false
}

func reverseSlice(slice []string) string {
  var temp []string
  for i := len(slice) - 1; i >= 0; i-- {
    temp = append(temp, slice[i])
  }
  return strings.Join(temp, " ")
}

func deleteElem(slice []string, target string) []string {
  var new []string
  for _, v := range slice {
    if v != target {
      new = append(new, v)
    }
  }
  return new
}

func processCommands(commandList []string, bSize int) {
  for _, v := range commandList {
    c := strings.Split(v, " ")
    if c[0] == "add" {
      key := polyHash(c[1], bSize)
      if containsString(key) && insideSlice(store[key], c[1]) {
        continue
      } else {
        store[key] = append(store[key], c[1])
      }
    } else if c[0] == "find" {
      key := polyHash(c[1], bSize)
      elem, exists := store[key]
      if exists {
        if insideSlice(elem, c[1]) {
          fmt.Println("yes")
        } else {
          fmt.Println("no")
        }
      } else {
        fmt.Println("no")
      }
    } else if c[0] == "check" {
      fmt.Println(reverseSlice(store[c[1]]))
    } else if c[0] == "del" {
      key := polyHash(c[1], bSize)
      _, exists := store[key]
      if exists {
        store[key] = deleteElem(store[key], c[1])
      } else {
        continue
      }
    }
  }
}

func main() {
  var commands []string
  var numCommands, bucketSize int
  fmt.Scanln(&bucketSize)
  fmt.Scanln(&numCommands)
  scanner := bufio.NewScanner(os.Stdin)

  for i := 0; i < numCommands; i++ {
    scanner.Scan()
    commands = append(commands, scanner.Text())
  }

  processCommands(commands, bucketSize)
}
