package websocket

import (
	"device-monitor-api/monitor"
	"time"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) retrieveAndPushData() {
	for {
		ds := []byte("Network")
		ds = append(ds, byte('\u0017'))
		dsBytes := monitor.GetDiskMonitorJSON()

		ds = append(ds, dsBytes...)
		ds = append(ds, byte('>'))
		h.broadcast <- ds

		bs := []byte("Disk")
		bs = append(bs, byte('\u0017'))
		bsBytes := monitor.GetNetMonitorJSON()

		bs = append(bs, bsBytes...)
		bs = append(bs, byte('>'))
		h.broadcast <- bs

		time.Sleep(500 * time.Millisecond)
	}
}

func (h *Hub) Run() {
	go h.retrieveAndPushData()
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
