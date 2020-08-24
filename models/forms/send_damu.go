package forms

type SendDanmuForm struct {
	Uid     uint32 `json:"uid" binding:"required"`
	Rid     uint32 `json:"rid" binding:"required"`
	Content string `json:"content" binding:"required"`
	Color   string `json:"color" binding:"required"`
}
