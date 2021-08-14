package controller

import (
	"encoding/json"
	"net/http"
)

type controller struct{}

func (c *controller) ReturnJson(w http.ResponseWriter, v interface{}) {
	json.NewEncoder(w).Encode(v)
}
