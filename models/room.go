package models

type Room struct {
	Rid  uint32 `xorm:"autoincr" json:"rid"`
	Uid  uint32 `json:"uid"`
	Name string `json:"name"`
}
