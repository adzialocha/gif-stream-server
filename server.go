package main

import (
  "github.com/subosito/gotenv"

  "log"
  "os"
  "net/http"
)

func init() {
  gotenv.Load()
}

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
