package controllers

import (
	"github.com/revel/revel"
)

// Login ログインユースケースのcontrollerです。
type Login struct {
	*revel.Controller
}

// Index 初期表示です。
func (c Login) Index() revel.Result {
	message := "[Login]"
	return c.Render(message)
}
