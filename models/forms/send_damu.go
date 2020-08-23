package forms

type SendDanmuForm struct {
	Uid     string `json:"uid" binding:"required"`
	Rid     string `json:"rid" binding:"required"`
	Content string `json:"content" binding:"required"`
	Color   string `json:"color" binding:"required"`
}
