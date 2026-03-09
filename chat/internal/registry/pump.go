package registry

func (c *Client) readPump(hub *Hub) {
    defer func() {
        hub.unregister <- c
        c.conn.Close()
    }()

    for {
        _, message, err := c.conn.ReadMessage()
        if err != nil {
            break
        }

        // For now, just broadcast incoming messages
        hub.broadcast <- message
    }
}

func (c *Client) writePump() {
    defer c.conn.Close()
    for msg := range c.send {
        err := c.conn.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            break
        }
    }
}
