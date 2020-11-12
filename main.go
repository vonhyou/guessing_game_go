package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  fmt.Println("<<GUESSING GAME>>")

  var guess, secretNumber, life int

  fmt.Print("Please select the difficulty [H]ard [M]edium [E]asy: ")

  const (
    easyRange, easyLife     int = 100, 7
    mediumRange, mediumLife int = 1000, 10
    hardRange, hardLife     int = 10000, 14
  )

  var difficulty string
  fmt.Scanln(&difficulty)

  seedr := rand.NewSource(time.Now().UnixNano())
  randr := rand.New(seedr)

  switch difficulty {
  case "E":
    secretNumber, life = randr.Intn(easyRange), easyLife
    fmt.Println("Range: 0 ~ 100   Life: 7")
  case "M":
    secretNumber, life = randr.Intn(mediumRange), mediumLife
    fmt.Println("Range: 0 ~ 1000  Life: 10")
  case "H":
    secretNumber, life = randr.Intn(hardLife), hardRange
    fmt.Println("Range: 0 ~ 10000 Life: 14")
  default:
    fmt.Println("Input Error!")
  }

  // fmt.Println(secretNumber)

  for life > 0 {
    fmt.Print("Input your number: ")
    fmt.Scanln(&guess)

    if guess < secretNumber {
      fmt.Println("TOO SMALL!")
    } else if guess > secretNumber {
      fmt.Println("TOO BIG")
    } else {
      fmt.Println("YOU WON!\nYour Score:", life*100)
      break
    }
    life--
    fmt.Println("Your life remain:", life)
  }
  if life == 0 {
    fmt.Println("YOU LOSE!")
  }
}
