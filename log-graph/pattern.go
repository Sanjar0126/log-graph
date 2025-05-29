package loggraph

import "regexp"

type Pattern struct {
	Name  string
	Regex *regexp.Regexp
	XAxis string
	YAxis string
	Color string
}

var patterns = []Pattern{
	{
		Name:  "query",
		Regex: regexp.MustCompile(`\[query\]:\s*([\d.]+)s`),
		XAxis: "Index",
		YAxis: "Query Time (s)",
		Color: "#ff6384",
	},
	{
		Name:  "query2",
		Regex: regexp.MustCompile(`\[query2\]:\s*([\d.]+)s`),
		XAxis: "Index",
		YAxis: "Query2 Time (s)",
		Color: "#36a2eb",
	},
}
