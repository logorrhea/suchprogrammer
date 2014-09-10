package controllers

import (
	"github.com/revel/revel"
	"io/ioutil"
	"net/http"
)

type Github struct {
	*revel.Controller
}

//https://api.github.com/search/repositories?q=suchprogrammer
func (c Github) Search() revel.Result {
	q := c.Params.Get("query")
	resp, err := http.Get("https://api.github.com/search/repositories?q=" + q)
	if err != nil {
		return c.RenderJson(err)
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		body := string(bodyBytes[:])
		return c.RenderJson(body)
	}
}
