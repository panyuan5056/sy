package task

import (
	"encoding/json"

	"sy/core"

	"sy/models"
	"sy/pkg/e"
	"sy/pkg/logging"

	"github.com/spf13/cast"
)

func Size() int64 {
	return models.Size()
}

func Pop() []TaskDetail {
	queues := models.Pop()
	details := []TaskDetail{}
	for _, queue := range queues {
		config := map[string]interface{}{}
		if err := json.Unmarshal([]byte(queue.Content), &config); err == nil {
			config["type"] = queue.Category
			config["id"] = cast.ToString(queue.ID)
			details = append(details, TaskDetail{fn: call, params: config})
		} else {
			logging.Error(err.Error())
		}
	}
	return details
}

func call(config map[string]interface{}) {
	report := core.WaterMark(cast.ToString(config["type"]), config)
	if content, err := json.Marshal(report); err == nil {
		models.Report(cast.ToString(config["id"]), string(content), e.GetMsg(e.SUCCESS), 3)
	} else {
		models.Report(cast.ToString(config["id"]), "", e.GetMsg(e.ERROR_FX), 4)
	}
}
