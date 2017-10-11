package controllers

import (
	"github.com/revel/revel"
)

// Sale 出品ユースケースcontrollerです。
type Sale struct {
	*revel.Controller
}

// Index 初期表示です。
func (c Sale) Index() revel.Result {
	message := "[Sale]"
	return c.Render(message)
}

// Confirm 確認画面です。
func (c Sale) Confirm() revel.Result {
	message := "[Sale]Confirm"
	return c.Render(message)
}

// Complete 出品処理です。
func (c Sale) Complete() revel.Result {

	return c.Redirect("/sale/completed")
}

// Completed 出品完了画面です。
func (c Sale) Completed() revel.Result {
	message := "[Sale]Complete"
	return c.Render(message)
}
