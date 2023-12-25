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
	"errors"
)

// PeopleApiService is a service that implents the logic for the PeopleApiServicer
// This service should implement the business logic for every endpoint for the PeopleApi API. 
// Include any external packages or services that will be required by this service.
type PeopleApiService struct {
}

// NewPeopleApiService creates a default api service
func NewPeopleApiService() PeopleApiServicer {
	return &PeopleApiService{}
}

// GetPeople - Return the people associated with a farm
func (s *PeopleApiService) GetPeople(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetPeople with the required logic for this service method.
	// Add api_people_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Person{}) or use other options such as http.Ok ...
	//return Response(200, []Person{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPeople method not implemented")
}

// GetPersonById - Find person by id
func (s *PeopleApiService) GetPersonById(ctx context.Context, personId int64) (ImplResponse, error) {
	// TODO - update GetPersonById with the required logic for this service method.
	// Add api_people_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Person{}) or use other options such as http.Ok ...
	//return Response(200, Person{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPersonById method not implemented")
}

