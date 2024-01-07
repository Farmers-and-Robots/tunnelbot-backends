// Code generated by openapi-generator-go DO NOT EDIT.
//
// Source:
//
//	Title: Tunnelbot - OpenAPI 3.0
//	Version: 1
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// GetEventByIdQueryParameters is an object.
type GetEventByIdQueryParameters struct {
	// EventId: id of event to return
	EventId int64 `json:"eventId" mapstructure:"eventId"`
}

// Validate implements basic validation for this model
func (m GetEventByIdQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetEventId returns the EventId property
func (m GetEventByIdQueryParameters) GetEventId() int64 {
	return m.EventId
}

// SetEventId sets the EventId property
func (m *GetEventByIdQueryParameters) SetEventId(val int64) {
	m.EventId = val
}
