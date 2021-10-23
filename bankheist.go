package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100) // Random value between 0 and 99

  // Conditional Statements
  if eludedGuards >= 50{
    fmt.Println("Step 1")
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  }else{
    isHeistOn = false
    fmt.Println("Step 1")
    fmt.Println("Plan a better disguise next time?")
  }
  openedVault := rand.Intn(100)

  // Second conditional statement
  if isHeistOn && openedVault >= 70{
    fmt.Println("Step 2")
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("Step 2")
    fmt.Println("The Vault cannot be opened.")
  }

  // Third conditional Statement
  leftSafely := rand.Intn(5)
  if isHeistOn{
    switch leftSafely{
      case 0:
        isHeistOn = false
        fmt.Println("Step 3")
        fmt.Println("We got caught with the Lasers...")
      case 1:
        isHeistOn = false
        fmt.Println("Step 3")
        fmt.Println("Turns out vault doors don't open from the inside...")
      case 2:
        isHeistOn = false
        fmt.Println("Step 3")
        fmt.Println("Turns out vault doors flooded from the inside...")
      case 3:
        isHeistOn = false
        fmt.Println("Step 3")
        fmt.Println("We got caught by bank Security...")
      default:
        isHeistOn = true
        fmt.Println("Step 3")
        fmt.Println("Start the getaway car!")
    }
  }
  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("$",amtStolen)
  }
  fmt.Println("")
  fmt.Println(isHeistOn)
}
