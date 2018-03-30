package controllers

import (
	"github.com/revel/revel"
)

// App class
type App struct {
	*revel.Controller
}

// Index method
func (c App) Index() revel.Result {
	return c.Render()
}
