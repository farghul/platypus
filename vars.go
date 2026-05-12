package main

type WPLatest struct {
	Offers []struct {
		Current string `json:"current"`
	} `json:"offers"`
}

// Changelogs builds a collection of urls to target changelogs
type Changelogs struct {
	ACFPro    string `json:"acfpro"`
	Gravity   string `json:"gravity"`
	Poly      string `json:"poly"`
	Spotlight string `json:"spotlight"`
	WordPress string `json:"wordpress"`
	WPExport  string `json:"wpexport"`
}

type Environment struct {
	Address   string `json:"address"`
	Install   string `json:"install"`
	Recipient string `json:"recipient"`
	Sender    string `json:"sender"`
	Server    string `json:"server"`
	User      string `json:"user"`
}

const (
	bv       string = "1.0.0"
	reset    string = "\033[0m"
	bgred    string = "\033[41m"
	green    string = "\033[32m"
	yellow   string = "\033[33m"
	bgyellow string = "\033[43m"
	halt     string = "program halted"
	meta     string = "/data/automation/jsons/"
	temp     string = "/data/automation/temp/"
)

var (
	version     WPLatest
	changelogs  Changelogs
	environment Environment
	jsons       = []string{meta + "changelogs.json", meta + "test.json"}
	remains     = []string{temp + "grepped.txt", temp + "temp.json", temp + "webscrape.txt"}
)
