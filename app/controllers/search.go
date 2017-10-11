package controllers

import (
	"github.com/revel/revel"
)

// Search /searchへのリクエストに関するcontrollerです。
type Search struct {
	*revel.Controller
}

// Index 初期表示です。
func (c Search) Index() revel.Result {
	message := "[Search]"
	return c.Render(message)
}

// Detail 詳細画面です。
func (c Search) Detail() revel.Result {
	itemID := c.Params.Route.Get("itemId")
	message := "[Search]Detail " + itemID
	return c.Render(message)
}
