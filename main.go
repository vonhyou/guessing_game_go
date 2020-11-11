package main

import "fmt"

func main() {
  fmt.Println("GUESSING GAME")

  var guess int
  var secretNumber, life = 42, 10

  for life > 0 {
    fmt.Print("Input your number:")
    fmt.Scanln(&guess)
    if guess < secretNumber {
      fmt.Println("TOO SMALL!")
    } else if guess > secretNumber {
      fmt.Println("TOO BIG")
    } else {
      fmt.Println("YOU WON!")
      break
    }
    life--
  }
  if life == 0 {
    fmt.Println("YOU LOSE!")
  }
}
