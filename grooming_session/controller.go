package grooming_session

import (
	"context"
	"errors"
	"github.com/dmmmd/scrumpoker/controller"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

// Routes

func Router(r chi.Router) {
	r.Get("/", actionIndex)
	r.Post("/", actionPost)

	r.Route("/{ID:[^/]+}", func(r chi.Router) {
		r.Use(controller.CreateItemContext(injectModel))
		r.Get("/", actionItem)
		//	r.Put("/", actionPut)
		//	r.Delete("/", actionDelete)
	})
}

// API resources

type groomingSessionResponse struct {
	*GroomingSession
}

func (rd *groomingSessionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func newResponse(model *GroomingSession) *groomingSessionResponse {
	return &groomingSessionResponse{GroomingSession: model}
}

func newListResponse(models []GroomingSession) []render.Renderer {
	var list []render.Renderer
	for i := range models {
		list = append(list, newResponse(&models[i]))
	}
	return list
}

type GroomingSessionRequest struct {
	*GroomingSession
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
	collection, _ := NewGroomingSessionStorage().NewCollection()

	var models []GroomingSession
	err := collection.Find().All(&models)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	controller.SendRestListResponse(w, r, newListResponse(models))
}

func actionItem(w http.ResponseWriter, r *http.Request) {
	model := r.Context().Value("model").(*GroomingSession)

	controller.SendRestItemResponse(w, r, newResponse(model))
}

func actionPost(w http.ResponseWriter, r *http.Request) {
	data := &GroomingSessionRequest{}
	if err := render.Bind(r, data); err != nil {
		controller.SendErrorResponse(w, r, controller.NewErrInvalidRequest(err))
		return
	}

	model, err := NewGroomingSessionStorage().Store(data.GroomingSession)
	if nil != err {
		controller.SendErrorResponse(w, r, controller.NewErrServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	controller.SendRestItemResponse(w, r, newResponse(model.(*GroomingSession)))
}

// Middleware

func injectModel(ctx context.Context, property string, id string) (context.Context, error) {
	model, err := NewGroomingSessionStorage().Load(id)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, property, model), nil
}
