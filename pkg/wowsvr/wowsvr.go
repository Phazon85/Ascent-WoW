package wowsvr

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//PostAccount handles the HTTP requests to create an account on a test wow server.
func PostAccount(acc, pass string) error {
	body, err := json.Marshal(map[string]string{
		"name":     acc,
		"password": pass,
	})
	if err != nil {
		return err
	}

	_, err = http.Post("127.0.0.1", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	return nil
}
