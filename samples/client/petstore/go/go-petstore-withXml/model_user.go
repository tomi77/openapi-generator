/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore
// User struct for User
type User struct {
	ID int64 `json:"id,omitempty" xml:"id"`
	Username string `json:"username,omitempty" xml:"username"`
	FirstName string `json:"firstName,omitempty" xml:"firstName"`
	LastName string `json:"lastName,omitempty" xml:"lastName"`
	Email string `json:"email,omitempty" xml:"email"`
	Password string `json:"password,omitempty" xml:"password"`
	Phone string `json:"phone,omitempty" xml:"phone"`
	// User Status
	UserStatus int32 `json:"userStatus,omitempty" xml:"userStatus"`
}
