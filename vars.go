package main

// Jira builds the Jira API address and update source
type Jira struct {
	Source string `json:"source"`
	Token  string `json:"token"`
	URL    string `json:"url"`
}

// Changelogs builds a collection of urls to target changelogs
type Changelogs struct {
	ACF       string `json:"acf"`
	Calendar  string `json:"calendar"`
	Gravity   string `json:"gravity"`
	Poly      string `json:"poly"`
	Spotlight string `json:"spotlight"`
	Tickets   string `json:"tickets"`
	Virtual   string `json:"virtual"`
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
	changelogs  Changelogs
	environment Environment
	remains     = []string{temp + "grepped.txt", temp + "temp.json", temp + "webscrape.txt"}
)
