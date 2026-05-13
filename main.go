package main

import (
	"fmt"
	"os"
)

// Launch the program and execute as directed by the supplied flag
func main() {
	var flag string = flags()
	logo()

	switch flag {
	case "-h", "--help":
		help()
	case "-r", "--run":
		credits()
		serialize()
		plugin()
	case "-v", "--version":
	case "--zero":
		alert("No flag detected - ")
	default:
		alert("Unknown argument(s) - ")
	}
}

// Test for an optional flag
func flags() string {
	var flag string

	if len(os.Args) == 1 {
		flag = "--zero"
	} else {
		flag = os.Args[1]
	}
	return flag
}

// Provide and highlight an informational message
func inform(message string) {
	Yellow.Printf("%s", "** ")
	fmt.Print(message)
	Yellow.Println(" **")
}

// Print a colourized error message
func alert(message string) {
	Red.Printf("\n%s", "Error: ")
	fmt.Printf("%s", message)
	BGRed.Println(halt)
	inform("Use -h to display help information")
	os.Exit(0)
}

// Print the help information
func help() {
	Yellow.Println("\nUsage:")
	fmt.Println("  [program] [flag]")
	Yellow.Println("\nOperational Flags:")
	Green.Printf("%s", "  -h, --help")
	fmt.Println("		Help Information")
	Green.Printf("%s", "  -r, --run")
	fmt.Println("		Run Program")
	Green.Printf("%s", "  -v, --version")
	fmt.Println("		Display Program Version")
	Yellow.Println("\nExample:")
	fmt.Println("  In your WordPress installation folder, run:")
	Green.Printf("%s", "    platypus -r")
	Yellow.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	Green.Println("    https://github.com/farghul/platypus.git")
}

func logo() {
	Magenta.Println("▗▄▄▖ ▗▖    ▗▄▖▗▄▄▄▖▗▖  ▗▖▗▄▄▖ ▗▖ ▗▖ ▗▄▄▖")
	Magenta.Println("▐▌ ▐▌▐▌   ▐▌ ▐▌ █   ▝▚▞▘ ▐▌ ▐▌▐▌ ▐▌▐▌   ")
	Magenta.Println("▐▛▀▘ ▐▌   ▐▛▀▜▌ █    ▐▌  ▐▛▀▘ ▐▌ ▐▌ ▝▀▚▖")
	Magenta.Println("▐▌   ▐▙▄▄▖▐▌ ▐▌ █    ▐▌  ▐▌   ▝▚▄▞▘▗▄▄▞▘")
	Magenta.Println(bv)
}

func credits() {
	fmt.Println("\nAn update search tool for WordPress")
	fmt.Println("Created by Byron Stuike")
}
