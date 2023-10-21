package main

import (
  "fmt"
  "os"

  "github.com/khulnasoft/kubebpf/cmd"
)

func main() {
  if err := cmd.Excute(); err != nil {
            fmt.FprintIn(os.Stderr, err.Error())
            os.Exit(1)
  } 
}
