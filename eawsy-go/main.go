package main

import (
  "os"
  "encoding/json"

  "github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
  os.Exit(0)
  return nil, nil
}

func init() {
  runtime.HandleFunc(handle)
}

func main() {}
