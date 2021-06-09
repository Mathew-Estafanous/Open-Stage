// Package docs Documentation for Open-Stage API
//
// The Open-Stage API is a REST service and is the platform for the
// live Q&A platform open stage. Allowing for the creation of rooms and
// associated questions within those rooms.
//
// Schemes: https
// BasePath: /v1/
// Host: open-stage-platform.herokuapp.com
// Version: 1.0
// License: MIT https://opensource.org/licenses/MIT
// Contact: Mathew Estafanous<mathewestafanous13@gmail.com> https://mathewestafanous.com/
//
// Security:
// - jwtAuth:
//
// SecurityDefinitions:
// jwtAuth:
//   type: apiKey
//   name: Authorization
//   in: header
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package docs

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/handler"
)

// A question that has been posted within a room.
// swagger:response questionResponse
type questionResponse struct {
	// in: body
	Body domain.Question
}

// A list of questions posted within a room.
// swagger:response multiQuestionResponse
type multiQuestionResponse struct {
	// in: body
	Body []domain.Question
}

// The conference room.
// swagger:response roomResponse
type roomResponse struct {
	// in: body
	Body domain.Room
}

// User Account information
// swagger:response accountResponse
type accountResponse struct {
	// in: body
	Body handler.AccountResp
}

// Both access and refresh tokens.
// swagger:response authToken
type tokenResponse struct {
	// in: body
	Body domain.AuthToken
}

// An http error response.
// swagger:response errorResponse
type errorResponse struct {
	// in: body
	Body handler.ResponseError
}

// swagger:parameters updateLikes
type updateLikesBody struct {
	// in: body
	Body handler.UpdateLike
}

// swagger:parameters createQuestion
type createQuestionBody struct {
	// in: body
	Body handler.NewQuestion
}

// swagger:parameters loginAccount
type loginBody struct {
	// in: body
	Body handler.Login
}

// swagger:parameters refresh
type refreshBody struct {
	// in: body
	Body handler.Refresh
}

// swagger:parameters createRoom
type createRoomBody struct {
	// in: body
	Body domain.Room
}

// swagger:parameters createAccount
type createAcountBody struct {
	// in: body
	Body handler.CreateAccount
}

// swagger:parameters roomCode
type roomCodePath struct {
	// The room code that all questions will be retrieved from.
	// in: path
	Code string `json:"roomCode"`
}

// swagger:parameters getCode
// swagger:parameters delCode
type codePath struct {
	// The unique room code.
	// in: path
	Code string `json:"code"`
}

// swagger:parameters questionId
type questionIdPath struct {
	// The question's ID
	// in: path
	ID string `json:"question_id"`
}

// swagger:parameters accountId
type accountIdPath struct {
	// The account ID
	// in: path
	ID string `json:"id"`
}
