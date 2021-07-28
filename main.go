package main

import (
	"fmt"
	"path"
	"unicode/utf8"
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
	// score := 0    //DO NOT DO THIS
	// var score int //DO THIS, SET TO 0 AUTOMATICALLY
	// // GREATER READABILITY EXAMPLE
	// var (
	// 	//RELATED:
	// 	video string
	// 	//CLOSELY RELATED
	// 	duration int
	// 	current  int
	// )

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

	// EXERCISE: Change color to blue
	color := "green"
	//color = "blue"
	fmt.Println("color: ", color)

	//EXERCISE: Variables To Variables
	//
	//  1. Change the value of `color` variable to "dark green"
	//
	//  2. Do not assign "dark green" to `color` directly
	//
	//     Instead, use the `color` variable again
	//     while assigning to `color`
	//
	//  3. Print it

	color = "dark " + color
	fmt.Println("color: ", color)

	// ---------------------------------------------------------
	// EXERCISE: Assign With Expressions
	//
	//  1. Multiply 3.14 with 2 and assign it to `n` variable
	//
	//  2. Print the `n` variable

	n := 0.
	n = 3.14 * 2
	fmt.Println(n)

	// EXERCISE: Find the Rectangle's Perimeter
	//
	//  1. Find the perimeter of a rectangle
	//     Its width is  5
	//     Its height is 6
	//
	//  2. Assign the result to the `perimeter` variable
	//
	//  3. Print the `perimeter` variable

	var (
		perimeter     int
		width, height = 5, 6
	)

	perimeter = 2 * (width + height)
	fmt.Println(perimeter)

	// EXERCISE: Multi Assign
	//
	//  1. Assign "go" to `lang` variable
	//     and assign 2 to `version` variable
	//     using a multiple assignment statement
	//
	//  2. Print the variables

	var (
		lang    string
		version int
	)

	lang, version = "go", 2

	fmt.Println(lang, "version", version)

	// EXERCISE: Multi Assign #2
	//
	//  1. Assign the correct values to the variables
	//     to match to the EXPECTED OUTPUT below
	//
	//  2. Print the variables
	// EXPECTED OUTPUT
	//  Air is good on Mars
	//  It's true
	//  It is 19.5 degrees
	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	// EXERCISE: Multi Short Func
	//
	// 	1. Declare two variables using multiple short declaration syntax
	//
	//  2. Initialize the variables using `multi` function below
	//
	//  3. Discard the 1st variable's value in the declaration *think of blank identifer*
	//
	//  4. Print only the 2nd variable
	_, b := multi()
	fmt.Println(b)

	// EXERCISE: Swapper
	//
	//  1. Change `color` to "orange"
	//     and `color2` to "green" at the same time
	//
	//     (use multiple-assignment)
	//
	//  2. Print the variables
	color1, color2 := "red", "blue"
	color1, color2 = "orange", "green"
	fmt.Println(color1, color2)

	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)

	// EXERCISE: Discard The File
	//
	//  1. Print only the directory using `path.Split`
	//
	//  2. Discard the file part

	dir, _ = path.Split("secret/file.txt")
	fmt.Println(dir)

	//Strings and Raw String Literals
	var s string
	s = "how are you?"
	s = `how are you?`
	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)
	s = `
	<html>
		<body> Hello </body>
	</html?`
	fmt.Println(s)

	fmt.Println("c:\\my\\dir\\file")
	fmt.Println(`c:\my\dir\file`)

	// Find String Length
	name := "carl"
	fmt.Println(len(name)) // len returns how many bytes in a string. ONLY USED FOR ENGLISH CHARACTERS.

	name = "González" // non-english characters can be 2 to 4 bytes
	fmt.Println(len(name))

	// Proper way to get length of a string when it is UTF8 encoded.
	fmt.Println(utf8.RuneCountInString(name)) // a rune can represent English and Non-English characters as well

	// msg := os.Args[1]
	// l := len(msg)
	// ex := strings.Repeat("!", l)
	// //banger := msg + strings.Repeat("!", l)
	// //banger = strings.ToUpper(banger)
	// banger := ex + msg + ex
	// fmt.Println(banger)

	// EXERCISE: Iota Months
	//
	//  1. Initialize the constants using iota.
	//  2. You should find the correct formula for iota.
	//
	// RESTRICTIONS
	//  1. Remove the initializer values from all constants.
	//  2. Then use iota once for initializing one of the
	//     constants.

	const (
		Nov = 11 - iota
		Oct
		Sep
	)

	fmt.Println(Sep, Oct, Nov)

	// EXERCISE: Iota Months #2
	//
	//  1. Initialize multiple constants using iota.
	//  2. Please follow the instructions inside the code.
	const _ = iota

	const (
		_ = iota
		Jan
		Feb
		Mar
	)

	// This should print: 1 2 3
	// Not: 0 1 2
	fmt.Println(Jan, Feb, Mar)

	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)

}

func multi() (int, int) {
	return 5, 4
}
