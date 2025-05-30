package loggraph

import (
	"fmt"
	"regexp"
)

type Pattern struct {
	Name  string
	Regex *regexp.Regexp
	XAxis string
	YAxis string
	Color string
}

func BuildPatterns(cfg *Config) ([]Pattern, error) {
	var patterns []Pattern
	for _, c := range cfg.Charts {
		re, err := regexp.Compile(c.Regex)
		if err != nil {
			return nil, fmt.Errorf("invalid regex for %s: %w", c.Name, err)
		}
		patterns = append(patterns, Pattern{
			Name:  c.Name,
			Regex: re,
			XAxis: c.XAxis,
			YAxis: c.YAxis,
			Color: c.Color,
		})
	}
	return patterns, nil
}
