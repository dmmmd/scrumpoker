package scrumpoker_models

import (
	"log"
)

const table string = "grooming_sessions"

type GroomingSession struct {
	AbstractModel
	Title string `db:"title" json:"title"`
}

func NewGroomingSession() *GroomingSession {
	return &GroomingSession{AbstractModel: *newAbstractModel(table)}
}

// Storage

type GroomingSessionStorage struct {
	Storage
}

func NewGroomingSessionStorage() *GroomingSessionStorage {
	storage := NewStorage(NewGroomingSession().GetTableName())
	return &GroomingSessionStorage{storage}
}

func (s GroomingSessionStorage) Load(id string) (*GroomingSession, error) {
	var m GroomingSession

	err := s.Storage.prepareLoadById(id).One(&m)
	if err != nil {
		log.Printf("Cannot load model: %q\n", err)
		return nil, err
	}

	return &m, nil
}
