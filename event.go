package stripe

import (
	"net/url"
)

// Event encapsulates details about a Event sent from Stripe
//
// see https://stripe.com/docs/api#events

// TODO: There is probably a better way to do this.
type EventData struct {
	Object             map[string]interface{} `json:"object"`
	PreviousAttributes map[string]interface{} `json:"previous_attributes"`
}

type Event struct {
	Id              string     `json:"id"`
	Object          string     `json:"object"`
	Data            *EventData `json:"data"`
	Livemode        bool       `json:"livemode"`
	Created         int64      `json:"created"`
	PendingWebhooks int64      `json:"pending_webhooks"`
	Type            string     `json:"type"`
	Request         string     `json:"request"`
}

type EventClient struct{}

// Retrieves a event with the given ID.
//
// see: https://stripe.com/docs/api#retrieve_event
func (self *EventClient) Retrieve(id string) (*Event, error) {
	event := Event{}
	path := "/v1/events/" + url.QueryEscape(id)
	err := query("GET", path, nil, &event)
	return &event, err
}

// Returns a list of the first 10 events.
//
// see: https://stripe.com/docs/api#list_events
func (self *EventClient) List() ([]*Event, error) {
	return self.ListWithFilters(Filters{})
}

// Returns a list of the events with all valid filters.
//
// see: https://stripe.com/docs/api#list_events
func (self *EventClient) ListWithFilters(filters Filters) ([]*Event, error) {
	// define a wrapper function for the Event List, so that we can
	// cleanly parse the JSON
	type listEventResp struct{ Data []*Event }
	resp := listEventResp{}

	values := url.Values{}
	addFiltersToValues([]string{"count", "offset", "type"}, filters, &values)

	err := query("GET", "/v1/events", values, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, err
}
