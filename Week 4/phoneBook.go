package main

import (
  "fmt"
  "strings"

  // For user input:
  "bufio"
  "os"
)

// Created to store the contacts AND their numbers in the form "number :
// contact"
var contacts = map[string]string {

}

// A helper function to check if "number" is in "contacts" or not. If it is,
// return the contact.  OTherwise, return "not found".
func isInsideContacts(number string) string {
  elem, ok := contacts[number]
  if ok {
    return elem
  } else {
    return "not found"
  }
}

// A function to solve the problem at hand:
func procesCommands(commandList []string) {
  for _, v := range commandList {
    c := strings.Split(v, " ")
    if c[0] == "add" {
      contacts[c[1]] = c[2]
    } else if c[0] == "del" {
      delete(contacts, c[1])
    } else if c[0] == "find" {
      fmt.Println(isInsideContacts(c[1]))
    }
  }
}

func main() {
  var commandList []string
  var numCommands int
  fmt.Scanln(&numCommands)
  scanner := bufio.NewScanner(os.Stdin)

  for i := 0; i < numCommands; i++ {
    scanner.Scan()
    commandList = append(commandList, scanner.Text())
  }

  procesCommands(commandList)
}
