package controllers

import (
	"github.com/go-chi/render"
	"net/http"
)

func SendRestListResponse(w http.ResponseWriter, r *http.Request, renderable []render.Renderer) {
	if err := render.RenderList(w, r, renderable); err != nil {
		_ = render.Render(w, r, errRender(err))
		return
	}
}

func SendRestItemResponse(w http.ResponseWriter, r *http.Request, renderable render.Renderer) {
	if err := render.Render(w, r, renderable); err != nil {
		_ = render.Render(w, r, errRender(err))
		return
	}
}

func SendErrorResponse(w http.ResponseWriter, r *http.Request, error *errResponse) {
	_ = render.Render(w, r, error)
}

type errResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func NewErrNotFound() *errResponse {
	return &errResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
}

func NewErrInvalidRequest(err error) *errResponse {
	return &errResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func NewErrServerError(err error) *errResponse {
	return &errResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Server error.",
		ErrorText:      err.Error(),
	}
}

func (e *errResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func errRender(err error) render.Renderer {
	return &errResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}
