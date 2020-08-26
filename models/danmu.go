package models

import "time"

type Danmu struct {
	Id      uint32    `xorm:"autoincr" json:"id"`
	Uid     uint32    `json:"uid"`
	Rid     uint32    `json:"rid"`
	Visible bool      `json:"visible"`
	Content string    `json:"content"`
	Color   string    `json:"color"`
	Created time.Time `json:"create_time"`
}
