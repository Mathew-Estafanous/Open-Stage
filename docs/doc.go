// Package docs Documentation for Open-Stage API
//
// The Open-Stage API is a REST service and is the platform for the
// live Q&A platform open stage. Allowing for the creation of rooms and
// associated questions within those rooms.
//
// Schemes: http https
// BasePath: /v1/
// Host: open-stage-platform.herokuapp.com
// Version: 1.0
// License: MIT http://opensource.org/licenses/MIT
// Contact: Mathew Estafanous<mathewestafanous13@gmail.com> https://mathewestafanous.com/
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

// An http error response.
// swagger:response errorResponse
type errorResponse struct {
	// in: body
	Body domain.ResponseError
}

// swagger:parameters updateLikes
type updateLikesReq struct {
	// in: body
	Body handler.UpdateLike
}
