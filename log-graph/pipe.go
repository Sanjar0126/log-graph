package loggraph

import (
	"bufio"
	"os"
	"strconv"
)

func HandleInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		seriesData := make(map[string]QueryData)

		for _, pat := range patterns {
			if match := pat.Regex.FindStringSubmatch(line); match != nil {
				val, _ := strconv.ParseFloat(match[1], 64)

				idx := indexCounter[pat.Name] + 1
				indexCounter[pat.Name] = idx

				entry := QueryData{Index: idx, Value: val}

				history[pat.Name] = append(history[pat.Name], entry)
				if len(history[pat.Name]) > historyLimit {
					history[pat.Name] = history[pat.Name][1:]
				}

				seriesData[pat.Name] = entry
			}
		}

		if len(seriesData) > 0 {
			broadcast <- BroadcastMessage{Type: "update", Data: seriesData}
		}
	}
}
