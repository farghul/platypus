package main

import (
	"os"
	"regexp"
	"strings"
)

// Run the functions to gather premium plugin versions currently installed and available
func subscription() string {
	var polylangInstalled = current("premium-plugin/polylang-pro")
	var exportInstalled = current("premium-plugin/wp-all-export-pro")
	var spotlightInstalled = current("freemius/spotlight-social-photo-feeds-premium")
	var polylangAvailable = latest(changelogs.Poly, "h4")
	var exportAvailable = latest(changelogs.WPExport, "h4")
	var spotlightAvailable = latest(changelogs.Spotlight, "h2")
	collect := results(polylangAvailable, polylangInstalled, "polylang-pro") + results(exportAvailable, exportInstalled, "wp-all-export-pro") + results(spotlightAvailable, spotlightInstalled, "spotlight-social-photo-feeds-premium")
	return collect
}

func wpcore() string {
	var coreInstalled = current("roots/wordpress")
	var coreAvailable = sift()
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
	where := strings.TrimSuffix(environment.Install, "web/wp") + "composer.json"
	what := concat("ssh", "-T", environment.User+"@"+environment.Server, " cat "+where)
	inspect(os.WriteFile(remains[1], what, 0666))
	grep, _ := capture("grep", p, remains[1])
	return regmatch(strings.TrimSpace(string(grep)))
}

// Find the latest versions of our premium plugins from the applicable websites
func latest(u, g string) string {
	capture("curl", "-s", u, "-o", remains[2])
	grep, _ := capture("grep", g, remains[2])
	inspect(os.WriteFile(remains[0], grep, 0666))
	head, _ := capture("head", "-n 1", remains[0])
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
