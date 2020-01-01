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
	"net/http"
	"os"
)


// PetAPIRouter defines the required methods for binding the api requests to a responses for the PetAPI
// The PetAPIRouter implementation should parse necessary information from the http request, 
// pass the data to a PetAPIServicer to perform the required actions, then write the service results to the http response.
type PetAPIRouter interface { 
	AddPet(http.ResponseWriter, *http.Request)
	DeletePet(http.ResponseWriter, *http.Request)
	FindPetsByStatus(http.ResponseWriter, *http.Request)
	FindPetsByTags(http.ResponseWriter, *http.Request)
	GetPetByID(http.ResponseWriter, *http.Request)
	UpdatePet(http.ResponseWriter, *http.Request)
	UpdatePetWithForm(http.ResponseWriter, *http.Request)
	UploadFile(http.ResponseWriter, *http.Request)
}
// StoreAPIRouter defines the required methods for binding the api requests to a responses for the StoreAPI
// The StoreAPIRouter implementation should parse necessary information from the http request, 
// pass the data to a StoreAPIServicer to perform the required actions, then write the service results to the http response.
type StoreAPIRouter interface { 
	DeleteOrder(http.ResponseWriter, *http.Request)
	GetInventory(http.ResponseWriter, *http.Request)
	GetOrderByID(http.ResponseWriter, *http.Request)
	PlaceOrder(http.ResponseWriter, *http.Request)
}
// UserAPIRouter defines the required methods for binding the api requests to a responses for the UserAPI
// The UserAPIRouter implementation should parse necessary information from the http request, 
// pass the data to a UserAPIServicer to perform the required actions, then write the service results to the http response.
type UserAPIRouter interface { 
	CreateUser(http.ResponseWriter, *http.Request)
	CreateUsersWithArrayInput(http.ResponseWriter, *http.Request)
	CreateUsersWithListInput(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	GetUserByName(http.ResponseWriter, *http.Request)
	LoginUser(http.ResponseWriter, *http.Request)
	LogoutUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)
}


// PetAPIServicer defines the api actions for the PetAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type PetAPIServicer interface { 
	AddPet(Pet) (interface{}, error)
	DeletePet(int64, string) (interface{}, error)
	FindPetsByStatus([]string) (interface{}, error)
	FindPetsByTags([]string) (interface{}, error)
	GetPetByID(int64) (interface{}, error)
	UpdatePet(Pet) (interface{}, error)
	UpdatePetWithForm(int64, string, string) (interface{}, error)
	UploadFile(int64, string, *os.File) (interface{}, error)
}


// StoreAPIServicer defines the api actions for the StoreAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type StoreAPIServicer interface { 
	DeleteOrder(string) (interface{}, error)
	GetInventory() (interface{}, error)
	GetOrderByID(int64) (interface{}, error)
	PlaceOrder(Order) (interface{}, error)
}


// UserAPIServicer defines the api actions for the UserAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type UserAPIServicer interface { 
	CreateUser(User) (interface{}, error)
	CreateUsersWithArrayInput([]User) (interface{}, error)
	CreateUsersWithListInput([]User) (interface{}, error)
	DeleteUser(string) (interface{}, error)
	GetUserByName(string) (interface{}, error)
	LoginUser(string, string) (interface{}, error)
	LogoutUser() (interface{}, error)
	UpdateUser(string, User) (interface{}, error)
}
