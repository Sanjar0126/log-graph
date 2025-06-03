package loggraph

import (
	"bufio"
	"os"
	"strconv"
)

func (h *WSHandler) HandleInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		seriesData := make(map[string]QueryData)

		Logger(INFO, line)

		for _, pat := range h.patterns {
			if match := pat.Regex.FindStringSubmatch(line); match != nil {
				val, _ := strconv.ParseFloat(match[1], 64)

				idx := h.indexCounter[pat.Name] + 1
				h.indexCounter[pat.Name] = idx

				entry := QueryData{Index: idx, Value: val}

				h.history[pat.Name] = append(h.history[pat.Name], entry)
				if len(h.history[pat.Name]) > h.historyLimit {
					h.history[pat.Name] = h.history[pat.Name][1:]
				}

				seriesData[pat.Name] = entry
			}
		}

		if len(seriesData) > 0 {
			h.broadcast <- BroadcastMessage{Type: "update", Data: seriesData}
		}
	}
}
