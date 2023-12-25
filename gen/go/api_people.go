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

// A PeopleApiController binds http requests to an api service and writes the service results to the http response
type PeopleApiController struct {
	service PeopleApiServicer
}

// NewPeopleApiController creates a default api controller
func NewPeopleApiController(s PeopleApiServicer) Router {
	return &PeopleApiController{ service: s }
}

// Routes returns all of the api route for the PeopleApiController
func (c *PeopleApiController) Routes() Routes {
	return Routes{ 
		{
			"GetPeople",
			strings.ToUpper("Get"),
			"/api/v1/people/",
			c.GetPeople,
		},
		{
			"GetPersonById",
			strings.ToUpper("Get"),
			"/api/v1/people/{personId}",
			c.GetPersonById,
		},
	}
}

// GetPeople - Return the people associated with a farm
func (c *PeopleApiController) GetPeople(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.GetPeople(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// GetPersonById - Find person by id
func (c *PeopleApiController) GetPersonById(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	personId, err := parseInt64Parameter(params["personId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.GetPersonById(r.Context(), personId)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
