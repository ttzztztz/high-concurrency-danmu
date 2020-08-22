package models

type Room struct {
	Id        uint32 `xorm:"pk autoincr"`
	CreatorId uint32 `xorm:"notnull"`
	Flow      int64  `xorm:"notnull default 0"`
}
