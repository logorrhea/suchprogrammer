package filters

import (
	"github.com/revel/revel"
	"strings"
)

type Action struct {
	Controller string
	Action     string
}

func ActionParseFilter(c *revel.Controller, filterChain []revel.Filter) {
	parsedAction := parseControllerAction(c.Action)
	c.RenderArgs["controller"] = parsedAction.Controller
	c.RenderArgs["action"] = parsedAction.Action
	filterChain[0](c, filterChain[1:])
}

func parseControllerAction(action string) Action {
	parts := make([]string, 0)
	for _, part := range strings.Split(action, ".") {
		parts = append(parts, strings.ToLower(part))
	}
	return Action{
		Controller: parts[0],
		Action:     parts[1],
	}
}
