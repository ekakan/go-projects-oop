package main

import "fmt"

func main(){
  // Defining the variables
  var publisher, writer, artist, title string
  var year, pageNumber int
  var grade float32

  // Assigning Values to the Comic book
  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5

  // Printing the comic book to the console
  // Mr. GoToSleep
  fmt.Println(title, "Written by:", writer, "Drawn by:", artist, "Published by:", publisher)
  fmt.Println("Number of pages:", pageNumber, "Publication year:", year, "Grade:", grade)
  fmt.Println("")
  // Assigning Values to the Comic book
  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  year = 2013
  pageNumber = 160
  grade = 9.0
  // Epic Vol. 1
  fmt.Println(title, "Written by:", writer, "Drawn by:", artist, "Published by:", publisher)
  fmt.Println("Number of pages:", pageNumber, "Publication year:", year, "Grade:", grade)
}
