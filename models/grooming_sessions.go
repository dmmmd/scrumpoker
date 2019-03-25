package scrumpoker_models

import (
	"github.com/dmmmd/scrumpoker/database"
	"github.com/google/uuid"
	"log"
	"upper.io/db.v3"
)

const TABLE string = "grooming_sessions"

type GroomingSession struct {
	//ID    int64  `db:"id,omitempty" json:"id,omitempty"`
	ID    uuid.UUID `db:"id,omitempty" json:"id,omitempty"`
	Title string    `db:"title" json:"title"`
}

// TODO Make sure to generate UUID on insert

func NewGroomingSessionsCollection() (db.Collection, error) {
	return database.NewCollection(TABLE)
}

func StoreGroomingSession(model *GroomingSession) (*GroomingSession, error) {
	// TODO improve UUID generation
	id, err := uuid.NewUUID()
	if nil != err {
		return nil, err
	}

	model.ID = id

	collection, _ := NewGroomingSessionsCollection()
	_, err = collection.Insert(model)
	if nil != err {
		return nil, err
	}

	return model, nil
}

func LoadGroomingSession(id string) (*GroomingSession, error) {
	collection, _ := NewGroomingSessionsCollection()

	var model GroomingSession
	err := collection.Find(db.Cond{"id": id}).One(&model)
	if err != nil {
		log.Printf("collection.One(): %q\n", err)
		return nil, err
	}

	return &model, nil
}
