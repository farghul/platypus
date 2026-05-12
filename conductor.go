package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
)

// Read the JSON files and Unmarshal the data into the appropriate Go structure
func serialize() {
	for index, element := range jsons {
		data, err := os.ReadFile(element)
		inspect(err)
		switch index {
		case 0:
			err := json.Unmarshal(data, &changelogs)
			inspect(err)
		case 1:
			err := json.Unmarshal(data, &environment)
			inspect(err)
		}
	}
}

// Trigger the search for updates
func plugin() {
	ups := wpcli("plugin", "list", "--update=available")
	gotcha(ups)
	premix := packagist(ups) + subscription() + wpcore()
	body := alphabetize(premix)
	if len(body) > 0 {
		err := os.WriteFile("/data/automation/lists/updates.txt", []byte(body+"\n"), 0666)
		inspect(err)
		mailing(body)
	} else {
		fmt.Println("No updates found for " + environment.Address)
	}
	for _, v := range remains {
		cleanup(v)
	}
}

// Filter the JSON blob response to get the latest version
func sift() string {
	data := get("https://api.wordpress.org/core/version-check/1.7/")
	err := json.Unmarshal(data, &version)
	inspect(err)
	if len(version.Offers) == 0 {
		log.Fatal("no offers")
	}
	current := version.Offers[0].Current
	return current
}

// Run the wp command to check for updates
func wpcli(x, y, z string) []string {
	c, _ := capture("wp", x, y, z, "--fields=name,version,update_version", "--format=csv", "--ssh="+environment.User+"@"+environment.Server+":"+environment.Install, "--url="+environment.Address, "--skip-plugins", "--skip-themes")
	f := strings.ReplaceAll(string(c), "\n", ",")
	r := strings.Split(f, ",")
	return r
}

// Format the output of plugin updates
func packagist(r []string) string {
	var value strings.Builder

	for a := 1; a < 4; a++ {
		r = slices.Delete(r, 0, 0+1)
	}

	for i := 0; i < len(r)-1; i++ {
		switch r[i] {
		case "gravityforms":
			value.WriteString("premium-plugin/" + r[i] + ":" + r[i+2] + "\n")
		default:
			value.WriteString("wpackagist-plugin/" + r[i] + ":" + r[i+2] + "\n")
		}
		i += 2
	}
	return strings.TrimRight(value.String(), " ")
}

// Alphabetize the update list before emailing it
func alphabetize(list string) string {
	s := strings.Split(list, "\n")
	sort.Strings(s)
	t := strings.Join(s, "\n")
	return t
}

// Catch any PHP errors which could interupt the program
func gotcha(output []string) {
	for i := range output {
		if strings.Contains(output[i], "Notice:") {
			alert("PHP Error on server interupting core functionality -")
		}
	}
}
