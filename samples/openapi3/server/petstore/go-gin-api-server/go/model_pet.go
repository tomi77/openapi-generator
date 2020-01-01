/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

type Pet struct {

	ID int64 `json:"id,omitempty"`

	Category Category `json:"category,omitempty"`

	Name string `json:"name"`

	PhotoURLs []string `json:"photoUrls"`

	Tags []Tag `json:"tags,omitempty"`

	// pet status in the store
	Status string `json:"status,omitempty"`
}
