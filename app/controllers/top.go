package controllers

import (
	"github.com/revel/revel"
)

// Top /リクエストに関するcontrollerです
type Top struct {
	*revel.Controller
}

// Index 初期表示です。
func (c Top) Index() revel.Result {
	message := "top"
	return c.Render(message)
}
