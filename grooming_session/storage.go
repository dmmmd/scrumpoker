package grooming_session

import (
	"github.com/dmmmd/scrumpoker/storage"
	"log"
)

type GroomingSessionStorage struct {
	storage.Storage
}

func NewGroomingSessionStorage() *GroomingSessionStorage {
	s := storage.NewStorage(NewGroomingSession().GetTableName())
	return &GroomingSessionStorage{Storage: s}
}

func (s GroomingSessionStorage) Load(id string) (*GroomingSession, error) {
	var m GroomingSession

	err := s.Storage.PrepareLoadById(id).One(&m)
	if err != nil {
		log.Printf("Cannot load model: %q\n", err)
		return nil, err
	}

	return &m, nil
}
