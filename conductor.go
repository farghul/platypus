package main

import (
	"os"
	"sort"
	"strings"
)

// Trigger the search for updates
func plugin() {
	short := []string{tmp, grp, web}
	if present() {
		ups := wpcli("plugin", "list", "--update=available")
		gotcha(ups)
		premix := packagist(ups) + assemble()
		body := alphabetize(premix)
		if len(body) > 0 {
			err := os.WriteFile(assets+"updates/updates.txt", []byte(body), 0666)
			inspect(err)
			mailman(body)
		} else {
			journal("No updates found for " + site)
		}
		for _, v := range short {
			cleanup(v)
		}
	}
}

// Run the wp command to check for updates
func wpcli(x, y, z string) []string {
	c := capture("wp", x, y, z, "--fields=name,version,update_version", "--format=csv", "--ssh="+user+":"+blog, "--url="+site)
	f := strings.ReplaceAll(string(c), "\n", ",")
	r := strings.Split(f, ",")
	return r
}

// Format the output of plugin updates
func packagist(r []string) string {
	var value string

	for a := 1; a < 4; a++ {
		r = append(r[:0], r[0+1:]...)
	}

	for i := 0; i < len(r)-1; i++ {
		switch r[i] {
		case "events-virtual":
			value += "premium-plugin/" + r[i] + ":" + r[i+2] + "\n"
		case "events-calendar-pro":
			value += "premium-plugin/" + r[i] + ":" + r[i+2] + "\n"
		case "gravityforms":
			value += "premium-plugin/" + r[i] + ":" + r[i+2] + "\n"
		default:
			value += "wpackagist-plugin/" + r[i] + ":" + r[i+2] + "\n"
		}
		i += 2
	}
	return strings.TrimRight(value, " ")
}

// Alphabetize the update list before emailing it
func alphabetize(list string) string {
	s := strings.Split(list, "\n")
	sort.Strings(s)
	t := strings.Join(s, "\n")
	return t
}

func gotcha(output []string) {
	for i := 0; i < len(output); i++ {
		if strings.Contains(output[i], "Notice:") {
			alert("PHP Error on server interupting core functionality -")
		}
	}
}
