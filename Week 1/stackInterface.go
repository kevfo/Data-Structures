package main

import (
  "fmt"
  "strings"
  "strconv"

  // For user input:
  "bufio"
  "os"
)

func processCommands(commandList []string) {
  var stack, maxes []int

  for _, v := range commandList {
    c := strings.Split(v, " ")
    if c[0] == "add" {
      convert, _ := strconv.Atoi(c[1])
      stack = append(stack, convert)
      if len(maxes) == 0 || maxes[len(maxes) - 1] < convert {
        maxes = append(maxes, convert)
      }
    } else if c[0] == "pop" {
      if stack[len(stack) - 1] == maxes[len(maxes) - 1] {
        maxes = maxes[:len(maxes) - 1]
      }
      stack = stack[:len(maxes) - 1]
    } else if c[0] == "max" {
      fmt.Println(maxes[len(maxes) - 1])
    }
  }
}

func main() {
  var commands []string
  var numCommands int
  fmt.Scanln(&numCommands)
  scanner := bufio.NewScanner(os.Stdin)

  for i := 0; i < numCommands; i++ {
    scanner.Scan()
    commands = append(commands, scanner.Text())
  }
}
