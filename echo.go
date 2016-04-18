package main

import (
  "fmt"
  "os"
  "strings"
  // "strconv"
)

func main()  {
// 1.1
  words := strings.Join(os.Args[0:], " ")

  fmt.Println(words)

// 1.2
  // for i := 0; i < len(os.Args); i++ {
  //   argWithIndex := strings.Join([]string{strconv.Itoa(i), os.Args[i]}, " ")
  //   // fmt.Println(i)
  //   // fmt.Println(os.Args[i])
  //   fmt.Println(argWithIndex)
  // }
}
