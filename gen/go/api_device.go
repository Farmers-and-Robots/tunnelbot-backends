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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DeviceApiController binds http requests to an api service and writes the service results to the http response
type DeviceApiController struct {
	service DeviceApiServicer
}

// NewDeviceApiController creates a default api controller
func NewDeviceApiController(s DeviceApiServicer) Router {
	return &DeviceApiController{ service: s }
}

// Routes returns all of the api route for the DeviceApiController
func (c *DeviceApiController) Routes() Routes {
	return Routes{ 
		{
			"AddDevice",
			strings.ToUpper("Post"),
			"/api/v1/device",
			c.AddDevice,
		},
		{
			"DeleteDevice",
			strings.ToUpper("Delete"),
			"/api/v1/device/{deviceId}",
			c.DeleteDevice,
		},
		{
			"GetDeviceById",
			strings.ToUpper("Get"),
			"/api/v1/device/{deviceId}",
			c.GetDeviceById,
		},
		{
			"GetDevices",
			strings.ToUpper("Get"),
			"/api/v1/device",
			c.GetDevices,
		},
		{
			"UpdateDevice",
			strings.ToUpper("Put"),
			"/api/v1/device",
			c.UpdateDevice,
		},
	}
}

// AddDevice - Add a new device to the db
func (c *DeviceApiController) AddDevice(w http.ResponseWriter, r *http.Request) { 
	device := &Device{}
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.AddDevice(r.Context(), *device)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// DeleteDevice - Deletes a device
func (c *DeviceApiController) DeleteDevice(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	deviceId, err := parseInt64Parameter(params["deviceId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	apiKey := r.Header.Get("api_key")
	result, err := c.service.DeleteDevice(r.Context(), deviceId, apiKey)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetDeviceById - Find device by id
func (c *DeviceApiController) GetDeviceById(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	deviceId, err := parseInt64Parameter(params["deviceId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.GetDeviceById(r.Context(), deviceId)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetDevices - Return the devices on a farm
func (c *DeviceApiController) GetDevices(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.GetDevices(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// UpdateDevice - Update an existing device
func (c *DeviceApiController) UpdateDevice(w http.ResponseWriter, r *http.Request) { 
	device := &Device{}
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	result, err := c.service.UpdateDevice(r.Context(), *device)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
