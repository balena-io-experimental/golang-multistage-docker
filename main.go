package main

import (
  "os"
  "fmt"
  "net/http"
)

const (
  port = ":80"
)

var calls = 0

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  calls++
  fmt.Fprintf(w, "Hello, world! You have called me %d times.\n", calls)
}

func main() {
  fmt.Printf("Started server on port %v.\n", port)
  fmt.Printf("Open https://%v.resindevice.io/ in your browser.\n", os.Getenv("RESIN_DEVICE_UUID"))
  http.HandleFunc("/", HelloWorld)
  http.ListenAndServe(port, nil)
}