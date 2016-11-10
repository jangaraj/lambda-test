package main

import (
  "os"

  "github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
  os.Exit(0)
}

func init() {
  runtime.HandleFunc(handle)
}

func main() {}
