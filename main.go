package getvero

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// GetVero keeps around the authtoken
type GetVero struct {
	AuthToken string
}

// Vero is an instance of vero
var Vero GetVero

type getveroResponse struct {
	Status  int
	Message string
}

func checkDataLength(eventWithDataAndExtras bool, a []interface{}) (int, error) {

	l := len(a)
	if l > 1 && !eventWithDataAndExtras {
		return l, errors.New("data interface can only be length of 0 or 1")
	} else if l != 2 && eventWithDataAndExtras {
		return l, errors.New("interface must be of length 2")
	}

	return l, nil

}

// sendToVero sends the request to getvero
func sendToVero(t string, d []byte, url string) error {

	req, err := http.NewRequest(t, url, bytes.NewBuffer(d))
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	c, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var r getveroResponse
	json.Unmarshal(c, &r)

	if r.Status != http.StatusOK {
		return errors.New(r.Message)
	}

	return nil

}

// sendsToVeroPost handles all the post requests to vero
func sendToVeroPost(d []byte, url string) error {

	return sendToVero("POST", d, url)

}

// sendToVeroPut handles all put requests to vero
func sendToVeroPut(d []byte, url string) error {

	return sendToVero("PUT", d, url)

}
