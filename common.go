package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// Remove files or directories
func cleanup(cut string) {
	inspect(os.Remove(cut))
}

// Run a terminal command, then capture and return the output as a byte
func capture(task string, args ...string) ([]byte, error) {
	return exec.Command(task, args...).CombinedOutput()
}

// Check for errors, halt the program if found, and log the result
func inspect(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// A generic GET request which captures the response
func get(target string) []byte {
	resp, err := http.Get(target)
	inspect(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	inspect(err)

	return body
}

// Run the Linux mail command and email the result to the configured recipent(s)
func mailing(list string) {
	cmd := exec.Command("mail", "-s", "WordPress updates for "+environment.Address, "-r", "Delivery Cactuar <"+environment.Sender+">", environment.Recipient)
	stdin, err := cmd.StdinPipe()
	inspect(err)

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, "Below is the current list of plugins requiring updates for "+environment.Address+". Have a magical day!\n\n"+list)
		inspect(err)
	}()

	out, _ := cmd.CombinedOutput() // Necessary although valueless

	log.Println(string(out))
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
