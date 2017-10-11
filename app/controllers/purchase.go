package controllers

import (
	"github.com/revel/revel"
)

// Purchase 購入ユースケースcontrollerです。
type Purchase struct {
	*revel.Controller
}

// Confirm 購入確認画面です。
func (c Purchase) Confirm() revel.Result {
	itemID := c.Params.Route.Get("itemId")
	message := "[Purchase]Confirm " + itemID
	return c.Render(message)
}

// Complete 購入処理です。
func (c Purchase) Complete() revel.Result {

	return c.Redirect("/purchase/completed")
}

// Completed 購入完了画面です。
func (c Purchase) Completed() revel.Result {
	message := "[Purchase]Completed"
	return c.Render(message)
}
