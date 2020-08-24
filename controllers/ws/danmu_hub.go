package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"strconv"
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
	clients    map[uint32]*Client
	Broadcast  chan *HubBroadcast
	register   chan *Client
	unregister chan *Client
}

func NewDanmuHub() *DanmuHub {
	danmuHub = &DanmuHub{
		clients:    make(map[uint32]*Client),
		Broadcast:  make(chan *HubBroadcast),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

	return danmuHub
}

func (h *DanmuHub) removeDanmuHubClient(uid uint32) {
	if client, ok := h.clients[uid]; ok {
		close(client.send)
		delete(h.clients, uid)
	}
}

func (h *DanmuHub) Run() {
	for {
		select {
		case client := <-h.register:
			if _, ok := h.clients[client.uid]; ok {
				close(h.clients[client.uid].send)
				delete(h.clients, client.uid)
			}

			h.clients[client.uid] = client
		case client := <-h.unregister:
			h.removeDanmuHubClient(client.uid)
		case broadcast := <-h.Broadcast:
			jsonByte, err := json.Marshal(HubBroadcastMessage{
				Content: broadcast.Content,
				Color:   broadcast.Color,
			})

			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, item := range h.clients {
				if item.rid == broadcast.Rid {
					item.send <- jsonByte
				}
			}
		}
	}
}
