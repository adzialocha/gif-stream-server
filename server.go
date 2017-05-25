package main

import (
  "log"
  "os"
  "net/http"
)

func main() {
  port := os.Getenv("PORT")

  if port == "" {
    log.Fatal("$PORT must be set")
  }

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)

  log.Println("Listening...")
  http.ListenAndServe(":" + port, nil)
}
