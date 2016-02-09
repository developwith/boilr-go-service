package main

import (
  "log"
  "fmt"
  "os"
  "runtime"
  "util"
  "service"
)

func main() {
  logger := log.New(os.Stderr, "{{AppName}}_wrapper", log.Lshortfile)
  setupMaxProcs()

  logger.Println("service:starting")
  if err := service.Run(os.Args[1:]); err != nil {
    fmt.Fprintf(os.Stderr, "Failed running server %q: %v\n", os.Args[1:], err)
    os.Exit(1)
  }
}

func setupMaxProcs() {
  numCPU := util.Round(float64(runtime.NumCPU()/2), 1, 0)
  runtime.GOMAXPROCS(int(numCPU))
}