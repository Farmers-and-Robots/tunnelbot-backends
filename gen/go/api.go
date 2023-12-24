/*
 * Tunnelbot - OpenAPI 3.0
 *
 * This is the API specification for the tunnelbot backend.
 *
 * API version: 1
 * Contact: info@farmersandrobots.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// DeviceApiRouter defines the required methods for binding the api requests to a responses for the DeviceApi
// The DeviceApiRouter implementation should parse necessary information from the http request, 
// pass the data to a DeviceApiServicer to perform the required actions, then write the service results to the http response.
type DeviceApiRouter interface { 
	AddDevice(http.ResponseWriter, *http.Request)
	DeleteDevice(http.ResponseWriter, *http.Request)
	GetDeviceById(http.ResponseWriter, *http.Request)
	UpdateDevice(http.ResponseWriter, *http.Request)
}
// FarmApiRouter defines the required methods for binding the api requests to a responses for the FarmApi
// The FarmApiRouter implementation should parse necessary information from the http request, 
// pass the data to a FarmApiServicer to perform the required actions, then write the service results to the http response.
type FarmApiRouter interface { 
	AddFarm(http.ResponseWriter, *http.Request)
	AddPerson(http.ResponseWriter, *http.Request)
	DeleteFarm(http.ResponseWriter, *http.Request)
	FindFarmByName(http.ResponseWriter, *http.Request)
	GetFarmById(http.ResponseWriter, *http.Request)
	RemovePerson(http.ResponseWriter, *http.Request)
	UpdateFarm(http.ResponseWriter, *http.Request)
}
// PeopleApiRouter defines the required methods for binding the api requests to a responses for the PeopleApi
// The PeopleApiRouter implementation should parse necessary information from the http request, 
// pass the data to a PeopleApiServicer to perform the required actions, then write the service results to the http response.
type PeopleApiRouter interface { 
	GetPersonById(http.ResponseWriter, *http.Request)
}
// TunnelApiRouter defines the required methods for binding the api requests to a responses for the TunnelApi
// The TunnelApiRouter implementation should parse necessary information from the http request, 
// pass the data to a TunnelApiServicer to perform the required actions, then write the service results to the http response.
type TunnelApiRouter interface { 
	AddTunnel(http.ResponseWriter, *http.Request)
	DeleteTunnel(http.ResponseWriter, *http.Request)
	GetTunnelById(http.ResponseWriter, *http.Request)
	UpdateTunnel(http.ResponseWriter, *http.Request)
}


// DeviceApiServicer defines the api actions for the DeviceApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type DeviceApiServicer interface { 
	AddDevice(context.Context, Device) (ImplResponse, error)
	DeleteDevice(context.Context, int64, string) (ImplResponse, error)
	GetDeviceById(context.Context, int64) (ImplResponse, error)
	UpdateDevice(context.Context, Device) (ImplResponse, error)
}


// FarmApiServicer defines the api actions for the FarmApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type FarmApiServicer interface { 
	AddFarm(context.Context, Farm) (ImplResponse, error)
	AddPerson(context.Context, int64, Farm) (ImplResponse, error)
	DeleteFarm(context.Context, int64, string) (ImplResponse, error)
	FindFarmByName(context.Context, string) (ImplResponse, error)
	GetFarmById(context.Context, int64) (ImplResponse, error)
	RemovePerson(context.Context, int64, Farm) (ImplResponse, error)
	UpdateFarm(context.Context, Farm) (ImplResponse, error)
}


// PeopleApiServicer defines the api actions for the PeopleApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type PeopleApiServicer interface { 
	GetPersonById(context.Context, int64) (ImplResponse, error)
}


// TunnelApiServicer defines the api actions for the TunnelApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type TunnelApiServicer interface { 
	AddTunnel(context.Context, Tunnel) (ImplResponse, error)
	DeleteTunnel(context.Context, int64, string) (ImplResponse, error)
	GetTunnelById(context.Context, int64) (ImplResponse, error)
	UpdateTunnel(context.Context, Tunnel) (ImplResponse, error)
}