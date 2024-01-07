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

// GetTunnelByIdQueryParameters is an object.
type GetTunnelByIdQueryParameters struct {
	// TunnelId: id of tunnel to return
	TunnelId int64 `json:"tunnelId" mapstructure:"tunnelId"`
}

// Validate implements basic validation for this model
func (m GetTunnelByIdQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetTunnelId returns the TunnelId property
func (m GetTunnelByIdQueryParameters) GetTunnelId() int64 {
	return m.TunnelId
}

// SetTunnelId sets the TunnelId property
func (m *GetTunnelByIdQueryParameters) SetTunnelId(val int64) {
	m.TunnelId = val
}