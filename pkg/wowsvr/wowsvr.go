package wowsvr

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//PostAccount handles the HTTP requests to create an account on a test wow server.
func PostAccount(acc, pass string) error {

	switch pass {
	case "":
		//assign default password
		body, err := json.Marshal(map[string]string{
			"name":     acc,
			"password": "password123",
		})
		if err != nil {
			return err
		}

		_, err = http.Post("mangos-account-creation:9000", "application/json", bytes.NewBuffer(body))
		if err != nil {
			return err
		}

	default:
		//asssign given password
		body, err := json.Marshal(map[string]string{
			"name":     acc,
			"password": pass,
		})
		if err != nil {
			return err
		}

		_, err = http.Post("mangos-account-creation:9000", "application/json", bytes.NewBuffer(body))
		if err != nil {
			return err
		}
	}

	return nil
}
