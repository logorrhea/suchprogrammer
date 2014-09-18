package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
	"net/http"
	"os/exec"
)

type Github struct {
	*revel.Controller
}

// These two structs are for unmarshalling the information
// that comes back from GitHub's API when requesting a file tree
// for a repository
type GHFile struct {
	Sha       string
	Url       string
	Tree      GHFileTree
	Truncated bool
}
type GHFileTree struct {
	Path string
	Mode string
	Type string
	Size int
	Sha  string
	Url  string
}

// These are for marshalling that same information, but in a different
// structure
type TreeNode struct {
	Name     string
	Parent   string
	Children []TreeNode
	Leaf     bool
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

	pullRepo(repo)

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

func (c Github) Tree() revel.Result {
	repo := c.Params.Get("repo")
	sha := c.Params.Get("sha")
	query := "https://api.github.com/repos/" + repo + "/git/trees/" + sha + "?recursive=1"
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

func (c Github) Testing() revel.Result {
	pullRepo(c.Params.Get("repo"))
	return c.RenderJson([]string{})
}

func pullRepo(repoName string) {

	// Change directory to user's temp directory

	// Run git pull {repoName}

	// Respond?

	execResults := exec.Command("git", "st")
	response, err := execResults.Output()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(response))
	}
	repoURL := "git@github.com:" + repoName
	fmt.Println(repoURL)
}

func parseFileTree(jsonData string) []string {

	// Parse the JSON into GHFile objects
	fileTree := make([]GHFile, 0)
	err := json.Unmarshal([]byte(jsonData), fileTree)
	if err != nil {
		revel.ERROR.Printf("%s", err.Error())
		return nil
	}

	// Loop through the GHFile array and start to construct a
	// directory tree
	//var dirTree map[string]
	//dirTree := []TreeNode{TreeNode{Name: "root", Parent: "", Children: make([]TreeNode, 0), Leaf: false}}
	//for _, ghFile := range fileTree {
	//h
	//}

	return nil
}
