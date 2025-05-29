package loggraph

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type QueryData struct {
	Index int     `json:"index"`
	Value float64 `json:"value"`
}

type ChartSeries struct {
	Name  string      `json:"name"`
	Data  []QueryData `json:"data"`
	XAxis string      `json:"xAxis"`
	YAxis string      `json:"yAxis"`
	Color string      `json:"color"`
}

type BroadcastMessage struct {
	Type string               `json:"type"`
	Data map[string]QueryData `json:"data"`
}

var (
	upgrader     = websocket.Upgrader{}
	clients      = make(map[*websocket.Conn]bool)
	mu           sync.Mutex
	broadcast    = make(chan BroadcastMessage)
	history      = make(map[string][]QueryData)
	historyLimit = 20
	indexCounter = make(map[string]int)
)

func getChartSeriesSnapshot() map[string]ChartSeries {
	charts := make(map[string]ChartSeries)
	for _, pat := range patterns {
		charts[pat.Name] = ChartSeries{
			Name:  pat.Name,
			Data:  history[pat.Name],
			XAxis: pat.XAxis,
			YAxis: pat.YAxis,
			Color: pat.Color,
		}
	}
	return charts
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer ws.Close()

	mu.Lock()
	clients[ws] = true

	if len(history) > 0 {
		ws.WriteJSON(struct {
			Type   string                 `json:"type"`
			Charts map[string]ChartSeries `json:"charts"`
		}{
			Type:   "init",
			Charts: getChartSeriesSnapshot(),
		})
	}
	mu.Unlock()

	for {
		if _, _, err := ws.NextReader(); err != nil {
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
	}
}

func HandleBroadcast() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("WebSocket error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
