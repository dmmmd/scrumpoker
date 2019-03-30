package grooming_session

import (
	"github.com/dmmmd/scrumpoker/storage"
)

const table string = "grooming_sessions"

type GroomingSession struct {
	storage.AbstractModel
	Title string `db:"title" json:"title"`
}

func NewGroomingSession() *GroomingSession {
	return &GroomingSession{AbstractModel: *storage.NewAbstractModel(table)}
}
