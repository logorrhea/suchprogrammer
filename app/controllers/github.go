package controllers

import (
	"fmt"
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

func (c Github) Commits() revel.Result {
	repo := c.Params.Get("repo")
	query := "https://api.github.com/repos/" + repo + "/commits"
	fmt.Println(query)
	resp, err := http.Get(query)

	var responseBody string
	if err != nil {
		responseBody = err.Error()
		c.Response.Status = 500
	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		responseBody = string(bodyBytes)
	}

	return c.RenderJson(responseBody)
}
