package main

import (
	"fmt"
	"path"
)

// Your customer wants you to record various measurements from a vehicles

// version := 0		//CANNOT DO THIS
// var version int 	//CAN DO THIS

func main() {
	// var speed int
	// fmt.Println(speed)

	var dir, file string
	dir, file = path.Split("css/main.css")

	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)

	var fileString string

	//Blank identifer
	_, fileString = path.Split("css/main.css")
	fmt.Println("file: ", fileString)

	//Short Declaration
	_, shortfile := path.Split("css/main.css")
	fmt.Println("short declaration file: ", shortfile)

	// WHEN TO USE NORMAL DECLARATION EXAMPLES BELOW
	score := 0    //DO NOT DO THIS
	var score int //DO THIS, SET TO 0 AUTOMATICALLY
	// GREATER READABILITY EXAMPLE
	var (
		//RELATED:
		video string
		//CLOSELY RELATED
		duration int
		current  int
	)

	// WHEN TO USE SHORT DECLARATION EXAMPLES BELOW
	// var width, height = 100, 50 // DO NOT DO THIS!
	// width, height := 100, 50 // DO THIS!

	// DO NOT DO THE FOLLOWING
	// width = 50 // assigns 50 to width
	// color := "red" // new variable: color
	// DO THIS
	// width , color := 50, "red" //change width to 50 and creates variable color=red

	// TYPE CONVERSION EXAMPLE BELOW
	speed := 100 // type int
	force := 2.5 // type float 64
	// speed = speed * force // Cannot do this b/c two diff types
	// speed = speed * int(force) // Produces wrong calculation
	// speed = float64(speed) * force // will not work b/c speed is still an int on the right hand side
	speed = int(float64(speed) * force)
	fmt.Println(speed)
	fmt.Println(force, int(force))

}
