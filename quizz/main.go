package main

import (
  "log"
  "os"
  "encoding/csv"
  "fmt"
  "io"
)

func main() {
  filePath := "problems.csv"
  f, err := os.Open(filePath)

  if err != nil {
    log.Fatal("Error while opening the file (%s): %s", filePath, err)
  }

  defer f.Close()

  reader := csv.NewReader(f)
  reader.Comma = ','
  reader.FieldsPerRecord = 2

  right := 0
  wrong := 0

  for {

    record, err := reader.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal("Error while reading the file (%s): %s", filePath, err)
    }

    q := record[0]
    a := record[1]
    
    fmt.Println(q, " ?")

    var name string
    fmt.Scan(&name)

    if name == a {
      right += 1
      fmt.Println("Correct!")
    } else {
      wrong += 1
      fmt.Println("Incorrect...")
    }
  }

  fmt.Printf("\nYou have a score of %d/%d !\n", right, right + wrong)
}
