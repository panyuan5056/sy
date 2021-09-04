package xdb

import (
	"encoding/json"
	"sy/pkg/logging"
)

type Config struct {
	Source  map[string]string `form:"source" json:"source" binding:"required"`
	Dump    map[string]string `form:"dump" json:"dump" binding:"required"`
	Content string            `form:"content" json:"dump" binding:"required"`
}

func (c *Config) str() (string, bool) {
	b, err := json.Marshal(c)
	if err != nil {
		logging.Error(err.Error())
		return "", false
	}
	return string(b), true
}
