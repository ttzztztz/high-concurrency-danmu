package models

import "time"

type Danmu struct {
	Id         uint32 `xorm:"pk autoincr"`
	UId        uint32 `xorm:"notnull comment('发送弹幕的用户id')"`
	RoomId     uint32 `xorm:"notnull comment('弹幕所属房间id')"`
	Visible    bool
	Content    string    `xorm:"notnull default '' ('弹幕内容')"`
	CreateTime time.Time `xorm:"created"`
}
