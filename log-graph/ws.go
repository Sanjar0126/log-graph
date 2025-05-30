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
// upgrader     = websocket.Upgrader{}
// clients      = make(map[*websocket.Conn]bool)
// mu           sync.Mutex
// broadcast    = make(chan BroadcastMessage)
// history      = make(map[string][]QueryData)
// historyLimit = 20
// indexCounter = make(map[string]int)
)

type WSHandler struct {
	patterns     []Pattern
	clients      map[*websocket.Conn]bool
	upgrader     websocket.Upgrader
	mu           sync.Mutex
	broadcast    chan BroadcastMessage
	history      map[string][]QueryData
	historyLimit int
	indexCounter map[string]int
}

func NewWSHandler(patterns []Pattern) *WSHandler {
	return &WSHandler{
		patterns:     patterns,
		clients:      make(map[*websocket.Conn]bool),
		upgrader:     websocket.Upgrader{},
		broadcast:    make(chan BroadcastMessage),
		history:      make(map[string][]QueryData),
		historyLimit: 20,
		indexCounter: make(map[string]int),
	}
}

func (h *WSHandler) getChartSeriesSnapshot() map[string]ChartSeries {
	charts := make(map[string]ChartSeries)
	for _, pat := range h.patterns {
		charts[pat.Name] = ChartSeries{
			Name:  pat.Name,
			Data:  h.history[pat.Name],
			XAxis: pat.XAxis,
			YAxis: pat.YAxis,
			Color: pat.Color,
		}
	}
	return charts
}

func (h *WSHandler) HandleConnections(w http.ResponseWriter, r *http.Request) {
	h.upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer ws.Close()

	h.mu.Lock()
	h.clients[ws] = true

	if len(h.history) > 0 {
		ws.WriteJSON(struct {
			Type   string                 `json:"type"`
			Charts map[string]ChartSeries `json:"charts"`
		}{
			Type:   "init",
			Charts: h.getChartSeriesSnapshot(),
		})
	}
	h.mu.Unlock()

	for {
		if _, _, err := ws.NextReader(); err != nil {
			h.mu.Lock()
			delete(h.clients, ws)
			h.mu.Unlock()
			break
		}
	}
}

func (h *WSHandler) HandleBroadcast() {
	for {
		msg := <-h.broadcast
		for client := range h.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("WebSocket error: %v", err)
				client.Close()
				delete(h.clients, client)
			}
		}
	}
}
