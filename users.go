package getvero

import (
	"encoding/json"
	"strings"
)

const usersURL string = "https://api.getvero.com/api/v2/users/"

type getveroIdentify struct {
	AuthToken string      `json:"auth_token"`
	ID        string      `json:"id"`
	Email     string      `json:"email,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type getveroAlias struct {
	AuthToken string `json:"auth_token"`
	ID        string `json:"id"`
	NewID     string `json:"new_id"`
}

type getveroUnsubscribe struct {
	AuthToken string `json:"auth_token"`
	ID        string `json:"id"`
}

type getveroResubscribe struct {
	AuthToken string `json:"auth_token"`
	ID        string `json:"id"`
}

type getveroDelete struct {
	AuthToken string `json:"auth_token"`
	ID        string `json:"id"`
}

// IdentifyUserWithEmail sends an identify request
func (v *GetVero) IdentifyUserWithEmail(id interface{}, e string, args ...interface{}) error {

	uID, err := checkID(id)
	if err != nil {
		return err
	}

	l, err := checkDataLength(false, args)
	if err != nil {
		return err
	}

	var d interface{}
	if l == 1 {
		d = args[0]
	}

	ident := getveroIdentify{
		AuthToken: v.AuthToken,
		ID:        uID,
		Email:     e,
		Data:      d,
	}
	b, err := json.Marshal(ident)

	if err != nil {
		return err
	}

	var url strings.Builder
	url.WriteString(usersURL)
	url.WriteString("track.json")

	err = sendToVeroPost(b, url.String())

	if err != nil {
		return err
	}

	return nil

}

// IdentifyUserWithoutEmail sends an identify request without the email
func (v *GetVero) IdentifyUserWithoutEmail(id interface{}, args ...interface{}) error {

	uID, err := checkID(id)
	if err != nil {
		return err
	}

	l, err := checkDataLength(false, args)
	if err != nil {
		return err
	}

	var d interface{}
	if l == 1 {
		d = args[0]
	}

	ident := &getveroIdentify{
		AuthToken: v.AuthToken,
		ID:        uID,
		Email:     "",
		Data:      d,
	}
	b, err := json.Marshal(ident)

	if err != nil {
		return err
	}

	var url strings.Builder
	url.WriteString(usersURL)
	url.WriteString("track.json")

	err = sendToVeroPost(b, url.String())

	if err != nil {
		return err
	}

	return nil

}

// Alias changes a users id
func (v *GetVero) Alias(id interface{}, newID interface{}) error {

	uID, err := checkID(id)
	if err != nil {
		return err
	}

	newUID, err := checkID(newID)
	if err != nil {
		return err
	}

	a := &getveroAlias{
		AuthToken: v.AuthToken,
		ID:        uID,
		NewID:     newUID,
	}
	b, err := json.Marshal(a)

	if err != nil {
		return err
	}

	var url strings.Builder
	url.WriteString(usersURL)
	url.WriteString("reidentify.json")

	err = sendToVeroPut(b, url.String())

	if err != nil {
		return err
	}

	return nil

}

// Unsubscribe removes user from email list
func (v *GetVero) Unsubscribe(id interface{}) error {

	uID, err := checkID(id)
	if err != nil {
		return err
	}

	unsub := &getveroUnsubscribe{
		AuthToken: v.AuthToken,
		ID:        uID,
	}
	b, err := json.Marshal(unsub)

	if err != nil {
		return err
	}

	var url strings.Builder
	url.WriteString(usersURL)
	url.WriteString("unsubscribe.json")

	err = sendToVeroPost(b, url.String())

	if err != nil {
		return err
	}

	return nil

}

// Resubscribe adds user back to mailing list
func (v *GetVero) Resubscribe(id interface{}) error {

	uID, err := checkID(id)
	if err != nil {
		return err
	}

	resub := &getveroResubscribe{
		AuthToken: v.AuthToken,
		ID:        uID,
	}
	b, err := json.Marshal(resub)

	if err != nil {
		return err
	}

	var url strings.Builder
	url.WriteString(usersURL)
	url.WriteString("resubscribe.json")

	err = sendToVeroPost(b, url.String())

	if err != nil {
		return err
	}

	return nil

}

// Delete removes user from getvero
func (v *GetVero) Delete(id interface{}) error {

	uID, err := checkID(id)
	if err != nil {
		return err
	}

	delete := &getveroDelete{
		AuthToken: v.AuthToken,
		ID:        uID,
	}
	b, err := json.Marshal(delete)

	if err != nil {
		return err
	}

	var url strings.Builder
	url.WriteString(usersURL)
	url.WriteString("delete.json")

	err = sendToVeroPost(b, url.String())

	if err != nil {
		return err
	}

	return nil

}
