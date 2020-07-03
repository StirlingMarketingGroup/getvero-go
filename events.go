package getvero

import (
	"encoding/json"
)

const eventsURL string = "https://api.getvero.com/api/v2/events/track.json"

type getveroEventIdentity struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type getveroEvent struct {
	AuthToken string               `json:"auth_token"`
	Identity  getveroEventIdentity `json:"identity"`
	EventName string               `json:"event_name"`
	Data      interface{}          `json:"data,omitempty"`
	Extras    interface{}          `json:"extras,omitempty"`
}

// TrackEvent tracks just the event name
func (v *GetVero) TrackEvent(id string, email string, event string) error {

	e := getveroEvent{
		AuthToken: v.AuthToken,
		Identity: getveroEventIdentity{
			ID:    id,
			Email: email,
		},
		EventName: event,
		Data:      nil,
		Extras:    nil,
	}
	b, err := json.Marshal(e)

	if err != nil {
		return err
	}

	err = sendToVeroPost(b, eventsURL)

	if err != nil {
		return err
	}

	return nil

}

// TrackEventWithData tracks the event with just event name and data
func (v *GetVero) TrackEventWithData(id string, email string, event string, args ...interface{}) error {

	l, err := checkDataLength(false, args)
	if err != nil {
		return err
	}

	var d interface{}
	if l == 1 {
		d = args[0]
	}

	e := getveroEvent{
		AuthToken: v.AuthToken,
		Identity: getveroEventIdentity{
			ID:    id,
			Email: email,
		},
		EventName: event,
		Data:      d,
		Extras:    nil,
	}
	b, err := json.Marshal(e)

	if err != nil {
		return err
	}

	err = sendToVeroPost(b, eventsURL)

	if err != nil {
		return err
	}

	return nil

}

// TrackEventWithExtra tracks event with just event name and extra
func (v *GetVero) TrackEventWithExtra(id string, email string, event string, args ...interface{}) error {

	l, err := checkDataLength(false, args)
	if err != nil {
		return err
	}

	var d interface{}
	if l == 1 {
		d = args[0]
	}

	e := getveroEvent{
		AuthToken: v.AuthToken,
		Identity: getveroEventIdentity{
			ID:    id,
			Email: email,
		},
		EventName: event,
		Data:      nil,
		Extras:    d,
	}
	b, err := json.Marshal(e)

	if err != nil {
		return err
	}

	err = sendToVeroPost(b, eventsURL)

	if err != nil {
		return err
	}

	return nil

}

// TrackEventWithDataAndExtra tracks event with event name, data, and extra
func (v *GetVero) TrackEventWithDataAndExtra(id string, email string, event string, args ...interface{}) error {

	_, err := checkDataLength(true, args)
	if err != nil {
		return err
	}

	data := args[0]
	extras := args[1]

	e := getveroEvent{
		AuthToken: v.AuthToken,
		Identity: getveroEventIdentity{
			ID:    id,
			Email: email,
		},
		EventName: event,
		Data:      data,
		Extras:    extras,
	}
	b, err := json.Marshal(e)

	if err != nil {
		return err
	}

	err = sendToVeroPost(b, eventsURL)

	if err != nil {
		return err
	}

	return nil

}
