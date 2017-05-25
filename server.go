package main

import (
  "github.com/joho/godotenv"

  "log"
  "os"
  "net/http"
)

func main() {
  godotenv.Load()

  port := os.Getenv("PORT")

  if port == "" {
    log.Fatal("$PORT must be set")
  }

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)

  log.Println("Listening...")
  http.ListenAndServe(":" + port, nil)
}
