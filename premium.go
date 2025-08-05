package main

import (
	"os"
	"regexp"
	"strings"
)

const (
	base string = "/data/automation/temp/"
	repo string = base + "platypus/"
	// Temp file locations
	web, tmp, grp string = base + "webscrape.txt", base + "temp.json", base + "grepped.txt"
)

// Run the functions to gather premium plugin versions currently installed and available
func subscription() string {
	var exportInstalled = current("premium-plugin/wp-all-export-pro")
	var ticketsInstalled = current("premium-plugin/event-tickets-plus")
	var polylangInstalled = current("premium-plugin/polylang-pro")
	var exportAvailable = latest(changelogs["wpexport"], "h4")
	var ticketsAvailable = latest(changelogs["tickets"], "Event Tickets Plus")
	var polylangAvailable = latest(changelogs["poly"], "h4")
	collect := results(ticketsAvailable, ticketsInstalled, "event-tickets-plus") + results(polylangAvailable, polylangInstalled, "polylang-pro") + results(exportAvailable, exportInstalled, "wp-all-export-pro")
	return collect
}

func wpcore() string {
	var coreInstalled = current("roots/wordpress")
	var coreAvailable = latest(changelogs["core"], "wp-block-wporg-release-version")
	collect := results(coreAvailable, coreInstalled, "wordpress")
	return collect
}

// Compare the version numbers and print the results if an update is available
func results(update, current, plugin string) string {
	var status string
	if update > current {
		if plugin == "wordpress" {
			status = "roots/" + plugin + ":" + update + "\n"
		} else {
			status = "premium-plugin/" + plugin + ":" + update + "\n"
		}
	}
	return status
}

// Find the current versions of our premium plugins from the composer.json file
func current(p string) string {
	where := strings.TrimSuffix(environment["install"], "web/wp") + "composer.json"
	what := concat("ssh", "-T", environment["user"]+"@"+environment["server"], " cat "+where)
	inspect(os.WriteFile(tmp, what, 0666))
	grep := capture("grep", p, tmp)
	return regmatch(strings.TrimSpace(string(grep)))
}

// Find the latest versions of our premium plugins from the applicable websites
func latest(u, g string) string {
	capture("curl", "-s", u, "-o", web)
	grep := capture("grep", g, web)
	inspect(os.WriteFile(grp, grep, 0666))
	head := capture("head", "-n 1", grp)
	return regmatch(strings.TrimSpace(string(head)))
}

// Remove all extraneous material, leaving only the version number itself
func regmatch(p string) string {
	var match []string
	tri := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}`)
	quad := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}.\d{1,3}`)

	if quad.MatchString(p) {
		match = quad.FindAllString(p, -1)
	} else if tri.MatchString(p) {
		match = tri.FindAllString(p, -1)
	}
	result := strings.Join(match, " ")
	return result
}
