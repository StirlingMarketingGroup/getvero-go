package getvero

import (
	"encoding/json"
)

const tagsURL string = "https://api.getvero.com/api/v2/users/tags/edit.json"

type getveroTag struct {
	AuthToken string   `json:"auth_token"`
	ID        string   `json:"id"`
	Add       []string `json:"add"`
	Remove    []string `json:"remove"`
}

// AddTags adds tags to a user
func (v *GetVero) AddTags(id string, tags []string) error {

	a := getveroTag{
		AuthToken: v.AuthToken,
		ID:        id,
		Add:       tags,
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
func (v *GetVero) RemoveTags(id string, tags []string) error {

	a := getveroTag{
		AuthToken: v.AuthToken,
		ID:        id,
		Add:       nil,
		Remove:    tags,
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
