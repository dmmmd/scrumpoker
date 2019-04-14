package api

import (
	"context"
	"github.com/dmmmd/scrumpoker/app/controller"
	"github.com/dmmmd/scrumpoker/app/grooming_session"
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
		r.Delete("/", actionDelete)
	})
}

// Actions

func actionIndex(w http.ResponseWriter, r *http.Request) {
	collection, _ := grooming_session.NewGroomingSessionStorage().NewCollection()

	var models []grooming_session.GroomingSession
	err := collection.Find().All(&models)
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}

	controller.SendRestListResponse(w, r, NewListResponse(models))
}

func actionItem(w http.ResponseWriter, r *http.Request) {
	model := r.Context().Value("model").(*grooming_session.GroomingSession)

	controller.SendRestItemResponse(w, r, NewItemResponse(model))
}

func actionPost(w http.ResponseWriter, r *http.Request) {
	data := &Request{}
	if err := render.Bind(r, data); err != nil {
		controller.SendErrorResponse(w, r, controller.NewErrInvalidRequest(err))
		return
	}

	model, err := grooming_session.NewGroomingSessionStorage().Store(data.GroomingSession)
	if nil != err {
		controller.SendErrorResponse(w, r, controller.NewErrServerError(err))
		return
	}

	render.Status(r, http.StatusCreated)
	controller.SendRestItemResponse(w, r, NewItemResponse(model.(*grooming_session.GroomingSession)))
}

func actionDelete(w http.ResponseWriter, r *http.Request) {
	model := r.Context().Value("model").(*grooming_session.GroomingSession)

	err := grooming_session.NewGroomingSessionStorage().Delete(model.GetId())
	if nil != err {
		controller.SendErrorResponse(w, r, controller.NewErrServerError(err))
		return
	}

	w.WriteHeader(204)
	controller.SendRawResponse(w, "")
}

// Middleware

func injectModel(ctx context.Context, property string, id string) (context.Context, error) {
	model, err := grooming_session.NewGroomingSessionStorage().Load(id)
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, property, model), nil
}
