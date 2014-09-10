package controllers

import (
	"github.com/revel/revel"
)

type Gource struct {
	*revel.Controller
}

func (c Gource) Index() revel.Result {
	c.RenderArgs["moreScripts"] = []string{"js/gource.js"}
	return c.Render()
}
