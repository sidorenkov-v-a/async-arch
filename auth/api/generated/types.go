// Package api_client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api_client

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Error defines model for Error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Token defines model for Token.
type Token struct {
	Token string `json:"token"`
}

// User defines model for User.
type User struct {
	Email     openapi_types.Email `json:"email"`
	FirstName string              `json:"firstName"`
	LastName  string              `json:"lastName"`
	Role      string              `json:"role"`
}

// UserLogin defines model for UserLogin.
type UserLogin struct {
	Email    openapi_types.Email `json:"email"`
	Password string              `json:"password"`
}

// UserRegister defines model for UserRegister.
type UserRegister struct {
	Email     openapi_types.Email `json:"email"`
	FirstName string              `json:"firstName"`
	LastName  string              `json:"lastName"`
	Password  string              `json:"password"`
	Role      string              `json:"role"`
}

// LoginUserJSONRequestBody defines body for LoginUser for application/json ContentType.
type LoginUserJSONRequestBody = UserLogin

// RegisterUserJSONRequestBody defines body for RegisterUser for application/json ContentType.
type RegisterUserJSONRequestBody = UserRegister