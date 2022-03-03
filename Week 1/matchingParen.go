package main

import (
  "fmt"
  "strings"
  "strconv"
)

// A helper function to check whether or not "x" is in "slice".  If it is, re-
// turn "true"; else, return false.
func isInside(slice []string, x string) bool {
  for _, v := range slice {
    if v == x {
      return true
    }
  }
  return false
}

// For the sake of convenience, define a struct to hold each entry:
type Entry struct {
  index int
  bracket string
}

// Solves the problem at hand:
func isBalanced(s string) string {
  var stack []Entry
  letters := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
  numAndSpecial := strings.Split("1234567890!@#$%^&*_+-=|\\;:<>?/.,~`", "")
  strList := strings.Split(strings.ToLower(s), "")
  opening := []string{"{", "[", "("}

  for i, v := range strList {
    if isInside(opening, v) {
      stack = append(stack, Entry{index: i, bracket: v})
    } else if isInside(letters, v) || isInside(numAndSpecial, v) {
      continue
    } else {
      if len(stack) == 0 {
        return strconv.Itoa(i + 1)
      }
      top := stack[len(stack) - 1]
      if (top.bracket == "{" && v != "}") || (top.bracket == "[" && v != "]") || (top.bracket == "(" && v != ")") {
        return strconv.Itoa(i + 1)
      } else {
        stack = stack[:len(stack) - 1]
      }
    }
  }
  if len(stack) == 0 {
    return "Success"
  } else {
    return strconv.Itoa(stack[len(stack) - 1].index + 1)
  }
}

func main() {
  var s string
  fmt.Scanln(&s)

  fmt.Println(isBalanced(s))
}
