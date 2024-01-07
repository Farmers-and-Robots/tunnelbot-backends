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

// Role is an enum.
type Role string

// Validate implements basic validation for this model
func (m Role) Validate() error {
	return InKnownRole.Validate(m)
}

var (
	RoleAdmin  Role = "admin"
	RoleOwner  Role = "owner"
	RoleWorker Role = "worker"

	// KnownRole is the list of valid Role
	KnownRole = []Role{
		RoleAdmin,
		RoleOwner,
		RoleWorker,
	}
	// KnownRoleString is the list of valid Role as string
	KnownRoleString = []string{
		string(RoleAdmin),
		string(RoleOwner),
		string(RoleWorker),
	}

	// InKnownRole is an ozzo-validator for Role
	InKnownRole = validation.In(
		RoleAdmin,
		RoleOwner,
		RoleWorker,
	)
)
