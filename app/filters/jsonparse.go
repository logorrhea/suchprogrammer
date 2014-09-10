package filters

import (
	"encoding/json"
	"github.com/revel/revel"
	"io/ioutil"
	"strings"
)

func JsonParamsFilter(c *revel.Controller, fc []revel.Filter) {
	if strings.Contains(c.Request.ContentType, "application/json") {
		data := map[string]string{}
		content, _ := ioutil.ReadAll(c.Request.Body)
		json.Unmarshal(content, &data)
		for k, v := range data {
			c.Params.Values.Set(k, v)
		}
	}
	fc[0](c, fc[1:])
}
