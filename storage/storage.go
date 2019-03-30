package storage

import (
	"github.com/dmmmd/scrumpoker/database"
	"github.com/google/uuid"
	"upper.io/db.v3"
)

type Storage struct {
	table string
}

func NewStorage(table string) Storage {
	return Storage{table: table}
}

func (s *Storage) NewCollection() (db.Collection, error) {
	return database.NewCollection(s.table)
}

func (s *Storage) Store(model Model) (Model, error) {
	id, err := s.generateId()
	if nil != err {
		return nil, err
	}

	model.SetId(id)

	collection, _ := s.NewCollection()
	_, err = collection.Insert(model)
	if nil != err {
		return nil, err
	}

	return model, nil
}

func (s *Storage) PrepareLoadById(id string) db.Result {
	collection, _ := s.NewCollection()

	return collection.Find(db.Cond{"id": id})
}

func (s *Storage) generateId() (uuid.UUID, error) {
	return uuid.NewUUID()
}
