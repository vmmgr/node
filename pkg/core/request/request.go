package request

import (
	"encoding/json"
	"github.com/vmmgr/node/pkg/core/client"
)

func SendServer(url, uuid string, progress uint, data string, error error) error {
	var comment string
	var status bool
	if error != nil {
		status = false
		comment = error.Error()
	} else {
		status = true
		comment = data
	}

	sendBody, _ := json.Marshal(Request{
		UUID:     uuid,
		Progress: progress,
		Status:   status,
		Comment:  comment,
	})
	err := client.Post(url, sendBody)
	return err
}
