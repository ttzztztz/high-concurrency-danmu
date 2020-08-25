package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
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
	defer func() {
		danmuHub.unregister <- c
		_ = c.conn.Close()
	}()

	c.conn.SetReadLimit(MaxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(PongWait))
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(PongWait))
		return nil
	})

	for {
		_, _, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(WriteWait))

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

func ServeDanmuWs(danmuHub *DanmuHub) func(*gin.Context) {
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
			hub:  danmuHub,
			conn: conn,
			send: make(chan []byte, 256),
			rid:  uint32(rid),
			uid:  uint32(uid),
		}
		danmuHub.register <- client

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
	Rooms      map[uint32]ConcurrentMap
	Broadcast  chan *HubBroadcast
	register   chan *Client
	unregister chan *Client

	mu sync.RWMutex
}

func NewDanmuHub() *DanmuHub {
	danmuHub = &DanmuHub{
		Rooms:      make(map[uint32]ConcurrentMap),
		Broadcast:  make(chan *HubBroadcast, 128),
		register:   make(chan *Client, 128),
		unregister: make(chan *Client, 128),
	}

	return danmuHub
}

func (h *DanmuHub) removeDanmuHubClient(client *Client) {
	h.mu.RLock()
	room := h.Rooms[client.rid]
	h.mu.RUnlock()

	room.Remove(client)
	if room.IsEmpty() {
		h.mu.Lock()
		if room.IsEmpty() {
			delete(h.Rooms, client.rid)
		}
		h.mu.Unlock()
	}
}

func (h *DanmuHub) sendBroadCast(broadcast *HubBroadcast) {
	jsonByte, err := json.Marshal(HubBroadcastMessage{
		Content: broadcast.Content,
		Color:   broadcast.Color,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	h.mu.RLock()
	room, ok := h.Rooms[broadcast.Rid]
	h.mu.RUnlock()

	if ok {
		room.IterCb(func(client *Client, _ interface{}) {
			client.send <- jsonByte
		})
	}
}

func (h *DanmuHub) joinClient(client *Client) {
	h.mu.RLock()
	room, ok := h.Rooms[client.rid]
	h.mu.RUnlock()

	if !ok {
		h.mu.Lock()
		if _room, ok := h.Rooms[client.rid]; !ok {
			room = MakeConcurrentHashmap()
			h.Rooms[client.rid] = room
		} else {
			room = _room
		}
		h.mu.Unlock()
	}

	room.Set(client, 1)
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
