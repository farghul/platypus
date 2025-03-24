package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"slices"
)

// Test if the server value passed to the program is on the list
func present() bool {
	return slices.Contains(servers, server)
}

// Run a terminal command, then capture and return the output as a byte
func capture(task string, args ...string) []byte {
	lpath, err := exec.LookPath(task)
	inspect(err)
	osCmd, _ := exec.Command(lpath, args...).CombinedOutput()
	return osCmd
}

// Check for errors, halt the program if found, and log the result
func inspect(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Run the Linux mail command and email the result to the configured recipent(s)
func mailman(list string) {
	cmd := exec.Command("mail", "-s", "WordPress updates for "+site, "-r", "Delivery Cactuar <"+sender+">", recipient)
	stdin, err := cmd.StdinPipe()
	inspect(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, "Below is the current list of plugins requiring updates for "+site+". Have a magical day!\n\n"+list)
		inspect(err)
	}()

	out, _ := cmd.CombinedOutput() // Necessary although valueless

	journal("Updates found and email sent" + string(out))
}

// Pipe together commands using the exec.Command function
func concat(method, flag, task, pipe string) []byte {
	cmd := exec.Command(method, flag, task)
	stdin, err := cmd.StdinPipe()
	inspect(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, pipe)
		inspect(err)
	}()

	out, _ := cmd.CombinedOutput()
	return out
}

// Record a message to a log file
func journal(message string) {
	file, err := os.OpenFile("/data/automation/logs/platypus.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	inspect(err)
	log.SetOutput(file)
	log.Println(message)
}

// Remove files or directories
func cleanup(cut string) {
	inspect(os.Remove(cut))
}
