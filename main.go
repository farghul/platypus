package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	bv       string = "1.0.0"
	reset    string = "\033[0m"
	bgred    string = "\033[41m"
	green    string = "\033[32m"
	yellow   string = "\033[33m"
	bgyellow string = "\033[43m"
	halt     string = "program halted"
)

var (
	environment map[string]string
	changelogs  map[string]string
)

// Launch the program and execute according to the results of the switch statement
func main() {

	json.Unmarshal([]byte(read(repo+"jsons/test.json")), &environment)
	json.Unmarshal([]byte(read(repo+"jsons/changelogs.json")), &changelogs)

	var flag string = flags()

	switch flag {
	case "-h", "--help":
		help()
	case "-r", "--run":
		plugin()
	case "-v", "--version":
		build()
	case "--zero":
		alert("No flag detected -")
	default:
		alert("Unknown argument(s) -")
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

// Print a colourized error message
func alert(message string) {
	fmt.Println("\n", bgred, message, halt, reset)
	fmt.Println(bgyellow, "Use -h for more detailed help information ")
	os.Exit(0)
}

// Display the build version of the program
func build() {
	fmt.Println("\n", yellow+"Platypus", green+bv, reset)
}

// Print the help information
func help() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  ./[program] [flag]")
	fmt.Println(yellow, "\nOperational Flags:")
	fmt.Println(green, " -h, --help", reset, "	   Help Information")
	fmt.Println(green, " -r, --run", reset, "      Run Program")
	fmt.Println(green, " -v, --version", reset, "  Display App Version")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("  In your WordPress installation folder, run:")
	fmt.Println(green, "    ./platypus")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "    https://github.com/farghul/platypus.git")
	fmt.Println(reset)
}
