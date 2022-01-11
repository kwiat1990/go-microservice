// Package classification Teams API
//
// Documentation for Teams API
//
//	Schemes: http
//	BasePath: /
//	Version: 0.1.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "go-microservice/data"

// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of teams
// swagger:response teamsResponse
type teamsResponseWrapper struct {
	// All current teams
	// in: body
	Body []data.Team
}

// Data structure representing a single team
// swagger:response teamResponse
type teamResponseWrapper struct {
	// Newly created team
	// in: body
	Body data.Team
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateTeam createTeam
type teamParamsWrapper struct {
	// team data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Team
}

// swagger:parameters getTeam deleteTeam
type teamIDParamsWrapper struct {
	// The id of the team for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
