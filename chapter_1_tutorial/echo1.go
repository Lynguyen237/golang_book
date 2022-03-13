package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	// print the name of the command itself
	fmt.Println("Name of command: ", os.Args[0])

	// measure how long the same for loop takes once put in a separate function
	// It's interesting that the for loop inside the main function takes longer when there's a function before it
	ForLoop()

	start := time.Now()

	var s, sep string

	// for initialization; condition; post
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	// print the string of arguments
	fmt.Println(s)

	duration := time.Since(start)
	fmt.Println("the for loop inside main() takes: ", duration)

	// The more efficient way with Join
	start2 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	duration2 := time.Since(start2)
	fmt.Println("join method takes: ", duration2)

	// if you run: "go run echo1.go a b c"
	// the expected output will be:
	// /var/folders/qv/.../exe/echo1 // name of the command or the path to the program
	// a b c

	// experiment with some other os methods
	fmt.Println("Getuid(): ", os.Getuid())
	fmt.Println("Geteuid():", os.Geteuid())
	fmt.Println(os.Getgroups())
	fmt.Printf("%v, %v", os.Getgroups())
	os.Exit(2)
}

func ForLoop() {
	defer Duration("ForLoop takes: ", time.Now())
	// start := time.Now()
	var s, sep string

	// for initialization; condition; post
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	// print the string of arguments
	fmt.Println(s)

	// duration := time.Since(start)
	// fmt.Println(duration)
}

func Duration(msg string, start time.Time) {
	log.Printf("%v %v\n", msg, time.Since(start))
}
