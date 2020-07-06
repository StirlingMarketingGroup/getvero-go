package getvero

import (
	"encoding/json"
	"errors"
)

const tagsURL string = "https://api.getvero.com/api/v2/users/tags/edit.json"

type getveroTag struct {
	AuthToken string   `json:"auth_token"`
	ID        string   `json:"id"`
	Add       []string `json:"add"`
	Remove    []string `json:"remove"`
}

func checkTags(tags ...interface{}) ([]string, error) {

	if len(tags) > 1 {
		t := make([]string, len(tags))
		for _, tag := range tags {
			if _, ok := tag.(string); ok {
				t = append(t, tag.(string))
			}
		}
		return t, nil
	}

	switch tags[0].(type) {
	case string:
		t := make([]string, 1)
		t[0] = tags[0].(string)
		return t, nil
	case []string:
		return tags[0].([]string), nil
	default:
		return nil, errors.New("tags must be of type string or []string")
	}

}

// AddTags adds tags to a user
func (v *GetVero) AddTags(id interface{}, tags ...interface{}) error {

	uID, err := checkID(id)
	if err != nil {
		return err
	}

	t, err := checkTags(tags...)
	if err != nil {
		return err
	}

	a := getveroTag{
		AuthToken: v.AuthToken,
		ID:        uID,
		Add:       t,
		Remove:    nil,
	}
	b, err := json.Marshal(a)

	if err != nil {
		return err
	}

	err = sendToVeroPut(b, tagsURL)

	if err != nil {
		return err
	}

	return nil

}

// RemoveTags removes tags to a user
func (v *GetVero) RemoveTags(id interface{}, tags ...interface{}) error {

	uID, err := checkID(id)
	if err != nil {
		return err
	}

	t, err := checkTags(tags...)
	if err != nil {
		return err
	}

	a := getveroTag{
		AuthToken: v.AuthToken,
		ID:        uID,
		Add:       nil,
		Remove:    t,
	}
	b, err := json.Marshal(a)

	if err != nil {
		return err
	}

	err = sendToVeroPut(b, tagsURL)

	if err != nil {
		return err
	}

	return nil

}
