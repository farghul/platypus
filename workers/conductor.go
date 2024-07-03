package workers

import (
	"log"
	"os"
)

// Plugin triggers the related functions
func Plugin() {
	short := []string{tmp, grp, web}
	if contains() {
		journal("Plugin update search triggered on " + site)
		ups := wpcli("plugin", "list", "--update=available")
		premix := freebies(ups) + assemble()
		body := alphabetize(premix)
		if len(body) > 0 {
			err := os.WriteFile("updates/updates.txt", []byte(body), 0666)
			inspect(err)
			mailman(body)
		} else {
			journal("No updates found")
		}
		for _, v := range short {
			cleanup(v)
		}
	}
}

// Theme triggers the related functions
func Theme() {
	journal("Theme update search triggered on " + site)
	/* TODO: Update themes */
}

// Core triggers the related functions
func Core() {
	journal("Core update search triggered on " + site)
	/* TODO: Update core components */
}

// Record a message to the log file
func journal(message string) {
	file, err := os.OpenFile("logs/platypus.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	inspect(err)
	log.SetOutput(file)
	log.Println(message)
}
