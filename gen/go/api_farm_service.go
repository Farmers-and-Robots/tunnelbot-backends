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

// FarmApiService is a service that implents the logic for the FarmApiServicer
// This service should implement the business logic for every endpoint for the FarmApi API. 
// Include any external packages or services that will be required by this service.
type FarmApiService struct {
}

// NewFarmApiService creates a default api service
func NewFarmApiService() FarmApiServicer {
	return &FarmApiService{}
}

// AddFarm - Add a new farm to the db
func (s *FarmApiService) AddFarm(ctx context.Context, farm Farm) (ImplResponse, error) {
	// TODO - update AddFarm with the required logic for this service method.
	// Add api_farm_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Farm{}) or use other options such as http.Ok ...
	//return Response(200, Farm{}), nil

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddFarm method not implemented")
}

// AddPerson - Add farm access for an existing person
func (s *FarmApiService) AddPerson(ctx context.Context, farmId int64, farm Farm) (ImplResponse, error) {
	// TODO - update AddPerson with the required logic for this service method.
	// Add api_farm_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Farm{}) or use other options such as http.Ok ...
	//return Response(200, Farm{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddPerson method not implemented")
}

// DeleteFarm - Deletes a farm
func (s *FarmApiService) DeleteFarm(ctx context.Context, farmId int64, apiKey string) (ImplResponse, error) {
	// TODO - update DeleteFarm with the required logic for this service method.
	// Add api_farm_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteFarm method not implemented")
}

// FindFarmByName - Finds a farm by name
func (s *FarmApiService) FindFarmByName(ctx context.Context, name string) (ImplResponse, error) {
	// TODO - update FindFarmByName with the required logic for this service method.
	// Add api_farm_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Farm{}) or use other options such as http.Ok ...
	//return Response(200, Farm{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("FindFarmByName method not implemented")
}

// GetFarmById - Find farm by id
func (s *FarmApiService) GetFarmById(ctx context.Context, farmId int64) (ImplResponse, error) {
	// TODO - update GetFarmById with the required logic for this service method.
	// Add api_farm_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Farm{}) or use other options such as http.Ok ...
	//return Response(200, Farm{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetFarmById method not implemented")
}

// RemovePerson - Remove farm access for an existing person
func (s *FarmApiService) RemovePerson(ctx context.Context, farmId int64, farm Farm) (ImplResponse, error) {
	// TODO - update RemovePerson with the required logic for this service method.
	// Add api_farm_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Farm{}) or use other options such as http.Ok ...
	//return Response(200, Farm{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("RemovePerson method not implemented")
}

// UpdateFarm - Update an existing farm
func (s *FarmApiService) UpdateFarm(ctx context.Context, farm Farm) (ImplResponse, error) {
	// TODO - update UpdateFarm with the required logic for this service method.
	// Add api_farm_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Farm{}) or use other options such as http.Ok ...
	//return Response(200, Farm{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateFarm method not implemented")
}

