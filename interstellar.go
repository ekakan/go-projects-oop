package main

import "fmt"

homePlanet := "Earth"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Printf("You have %v litre(s) of fuel left", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet{
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    default:
      fuel = 0
  }
  fmt.Println(fuel)
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string){
  fmt.Printf("Welcome to %v ! \n", planet)
}

// Create the function cantFly() here
func cantFly(){
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int, currentPlanet string
  currentPlanet = planet
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  greetPlanet(planet)
  fuelRemaining -= fuelCost
  if fuelCost > fuelRemaining {
    cantFly()
  }
  fmt.Printf("Fuel Remaining: %v litre(s)\n", fuelRemaining)
  return fuelRemaining
}

func goHome(planet string){
  fmt.Printf("Going back home to %v", planet)
}

func main() {
  // Test your functions!
  //fuelGauge(12)
  //calculateFuel("Mars")
  //greetPlanet("Earth")
  //cantFly()
  //flyToPlanet("Mars", 800000)

  // Create `planetChoice` and `fuel`
  fuel := 1000000
  planetChoice := "Venus"
  flyToPlanet(planetChoice, fuel)
  fuelGauge(flyToPlanet(planetChoice, fuel))
  // could also assign the function to the variable fuel

  // And then liftoff!

}
