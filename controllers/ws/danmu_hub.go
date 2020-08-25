package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"time"
)

var (
	danmuHub *DanmuHub
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	uid  uint32
	rid  uint32
}

func (c *Client) readPump() {
	c.conn.SetReadLimit(MaxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(PongWait)); err != nil {
		fmt.Println(err)
	}
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(PongWait))
		return nil
	})
}

func (c *Client) writePump() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if err := c.conn.SetWriteDeadline(time.Now().Add(WriteWait)); err != nil {
				fmt.Println(err)
			}

			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					fmt.Println(err)
				}
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Println(err)
				return
			}

			if _, err := w.Write(message); err != nil {
				fmt.Println(err)
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				if _, err := w.Write(<-c.send); err != nil {
					fmt.Println(err)
				}
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.conn.SetWriteDeadline(time.Now().Add(WriteWait)); err != nil {
				fmt.Println(err)
				return
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func ServeDanmuWs(contestHub *DanmuHub) func(*gin.Context) {
	return func(c *gin.Context) {
		_rid, _uid := c.Param("rid"), c.Param("uid")

		rid, err := strconv.ParseUint(_rid, 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}

		uid, err := strconv.ParseUint(_uid, 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}

		conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &Client{
			conn: conn,
			send: make(chan []byte, 256),
			rid:  uint32(rid),
			uid:  uint32(uid),
		}
		contestHub.register <- client

		go client.writePump()
		go client.readPump()
	}
}

type HubBroadcast struct {
	Content string `json:"content"`
	Color   string `json:"color"`
	Uid     uint32 `json:"uid"`
	Rid     uint32 `json:"rid"`
}

type HubBroadcastMessage struct {
	Content string `json:"content"`
	Color   string `json:"color"`
}

type DanmuHub struct {
	clients map[*Client]bool
	mu      sync.RWMutex

	rooms      map[uint32]map[*Client]bool
	Broadcast  chan *HubBroadcast
	register   chan *Client
	unregister chan *Client
}

func NewDanmuHub() *DanmuHub {
	danmuHub = &DanmuHub{
		clients:    make(map[*Client]bool),
		rooms:      make(map[uint32]map[*Client]bool),
		Broadcast:  make(chan *HubBroadcast),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	return danmuHub
}

func (h *DanmuHub) removeDanmuHubClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clients, client)
	delete(h.rooms[client.rid], client)
	if len(h.rooms[client.rid]) == 0 {
		delete(h.rooms, client.rid)
	}
}

func (h *DanmuHub) sendBroadCast(broadcast *HubBroadcast) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	jsonByte, err := json.Marshal(HubBroadcastMessage{
		Content: broadcast.Content,
		Color:   broadcast.Color,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	if _, ok := h.rooms[broadcast.Rid]; ok {
		for client := range h.rooms[broadcast.Rid] {
			client.send <- jsonByte
		}
	}
}

func (h *DanmuHub) joinClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[client] = true
	if _, ok := h.rooms[client.rid]; !ok {
		h.rooms[client.rid] = make(map[*Client]bool)
	}
	h.rooms[client.rid][client] = true
}

func (h *DanmuHub) Run() {
	for {
		select {
		case client := <-h.register:
			go h.joinClient(client)
		case client := <-h.unregister:
			go h.removeDanmuHubClient(client)
		case broadcast := <-h.Broadcast:
			go h.sendBroadCast(broadcast)
		}
	}
}
