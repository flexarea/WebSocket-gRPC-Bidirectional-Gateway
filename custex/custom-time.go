package custex

import "time"

type Client struct{}

func (c *Client) GetNano() int64 {
	return time.Now().UnixNano()
}

func (c *Client) GetMicro() int64 {
	return time.Now().UnixMicro()
}

func (c *Client) GetMilli() int64 {
	return time.Now().UnixMilli()
}
