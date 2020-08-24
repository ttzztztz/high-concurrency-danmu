package models

import "time"

type Danmu struct {
	Id         uint32 `xorm:"pk autoincr"`
	UId        uint32 `xorm:"notnull comment('发送弹幕的用户id')"`
	RoomId     uint32 `xorm:"notnull comment('弹幕所属房间id')"`
	Visible    bool
	Content    string    `xorm:"notnull default '' comment('弹幕内容')"`
	Color      string    `xorm:"notnull default 'black' comment('弹幕颜色')""`
	CreateTime time.Time `xorm:"created"`
}
