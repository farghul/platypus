package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

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
	cmd := exec.Command("mail", "-s", "WordPress updates for "+environment["address"], "-r", "Delivery Cactuar <"+environment["sender"]+">", environment["recipient"])
	stdin, err := cmd.StdinPipe()
	inspect(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, "Below is the current list of plugins requiring updates for "+environment["address"]+". Have a magical day!\n\n"+list)
		inspect(err)
	}()
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

// Read any file and return the contents as a byte variable
func read(file string) []byte {
	mission, err := os.Open(file)
	inspect(err)
	outcome, err := io.ReadAll(mission)
	inspect(err)
	defer mission.Close()
	return outcome
}

// Remove files or directories
func cleanup(cut string) {
	inspect(os.Remove(cut))
}
