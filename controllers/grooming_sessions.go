package controllers

import (
	"context"
	"errors"
	"github.com/dmmmd/scrumpoker/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

// Routes

func GroomingSessionsRouter(r chi.Router) {
	r.Get("/", actionIndex)
	r.Post("/", actionPost)

	r.Route("/{ID:[^/]+}", func(r chi.Router) {
		r.Use(itemContext)
		r.Get("/", actionItem)
		//	r.Put("/", actionPut)
		//	r.Delete("/", actionDelete)
	})
}

// API resources

type groomingSessionResponse struct {
	*scrumpoker_models.GroomingSession
}

func (rd *groomingSessionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newResponse(model *scrumpoker_models.GroomingSession) *groomingSessionResponse {
	return &groomingSessionResponse{GroomingSession: model}
}

func newListResponse(models []scrumpoker_models.GroomingSession) []render.Renderer {
	var list []render.Renderer
	for i := range models {
		list = append(list, newResponse(&models[i]))
	}
	return list
}

type GroomingSessionRequest struct {
	*scrumpoker_models.GroomingSession
}

func (post *GroomingSessionRequest) Bind(r *http.Request) error {
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

// Actions

func actionIndex(w http.ResponseWriter, r *http.Request) {
	collection, _ := scrumpoker_models.NewGroomingSessionStorage().NewCollection()

	var models []scrumpoker_models.GroomingSession
	err := collection.Find().All(&models)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	SendRestListResponse(w, r, newListResponse(models))
}

func actionItem(w http.ResponseWriter, r *http.Request) {
	model := r.Context().Value("model").(*scrumpoker_models.GroomingSession)

	SendRestItemResponse(w, r, newResponse(model))
}

func actionPost(w http.ResponseWriter, r *http.Request) {
	data := &GroomingSessionRequest{}
	if err := render.Bind(r, data); err != nil {
		SendErrorResponse(w, r, NewErrInvalidRequest(err))
		return
	}

	model, err := scrumpoker_models.NewGroomingSessionStorage().Store(data.GroomingSession)
	if nil != err {
		SendErrorResponse(w, r, NewErrServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	SendRestItemResponse(w, r, newResponse(model.(*scrumpoker_models.GroomingSession)))
}

// Middleware

// itemContext middleware is used to load a model object from
// the URL parameters passed through as the request. In case
// the model could not be found, we stop here and return a 404.
func itemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var model *scrumpoker_models.GroomingSession
		var err error

		if id := chi.URLParam(r, "ID"); id != "" {
			model, err = scrumpoker_models.NewGroomingSessionStorage().Load(id)
		} else {
			SendErrorResponse(w, r, NewErrNotFound())
			return
		}

		if err != nil {
			SendErrorResponse(w, r, NewErrNotFound())
			return
		}

		ctx := context.WithValue(r.Context(), "model", model)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
