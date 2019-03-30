package api

import (
	"errors"
	"github.com/dmmmd/scrumpoker/grooming_session"
	"github.com/go-chi/render"
	"net/http"
)

// API Request

type Request struct {
	*grooming_session.GroomingSession
}

func (post *Request) Bind(r *http.Request) error {
	// TODO properly ignore ID on POST
	// Now it's:
	// curl -d '{"id": "Umad?", "title":"Session E"}' -H "Content-Type: application/json" -X POST http://localhost:9090/grooming_sessions
	// {"status":"Invalid request.","error":"invalid UUID length: 5"}

	// post.GroomingSession is nil if no GroomingSession fields are sent in the request. Return an
	// error to avoid post nil pointer dereference.
	if post.GroomingSession == nil {
		return errors.New("missing required grooming session fields")
	}

	return nil
}

// API Response

type Response struct {
	*grooming_session.GroomingSession
}

func NewItemResponse(model *grooming_session.GroomingSession) *Response {
	return &Response{GroomingSession: model}
}

func NewListResponse(models []grooming_session.GroomingSession) []render.Renderer {
	var list []render.Renderer
	for i := range models {
		list = append(list, NewItemResponse(&models[i]))
	}
	return list
}

func (*Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
