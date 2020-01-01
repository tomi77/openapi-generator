/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"errors"
)

// StoreAPIService is a service that implents the logic for the StoreAPIServicer
// This service should implement the business logic for every endpoint for the StoreAPI API. 
// Include any external packages or services that will be required by this service.
type StoreAPIService struct {
}

// NewStoreAPIService creates a default api service
func NewStoreAPIService() StoreAPIServicer {
	return &StoreAPIService{}
}

// DeleteOrder - Delete purchase order by ID
func (s *StoreAPIService) DeleteOrder(orderID string) (interface{}, error) {
	// TODO - update DeleteOrder with the required logic for this service method.
	// Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'DeleteOrder' not implemented")
}

// GetInventory - Returns pet inventories by status
func (s *StoreAPIService) GetInventory() (interface{}, error) {
	// TODO - update GetInventory with the required logic for this service method.
	// Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetInventory' not implemented")
}

// GetOrderByID - Find purchase order by ID
func (s *StoreAPIService) GetOrderByID(orderID int64) (interface{}, error) {
	// TODO - update GetOrderByID with the required logic for this service method.
	// Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetOrderByID' not implemented")
}

// PlaceOrder - Place an order for a pet
func (s *StoreAPIService) PlaceOrder(body Order) (interface{}, error) {
	// TODO - update PlaceOrder with the required logic for this service method.
	// Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PlaceOrder' not implemented")
}
