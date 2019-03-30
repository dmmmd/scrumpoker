package storage

import "github.com/google/uuid"

type Model interface {
	GetTableName() string
	GetId() uuid.UUID
	SetId(id uuid.UUID)
}

type AbstractModel struct {
	table string
	ID    uuid.UUID `db:"id,omitempty" json:"id,omitempty"`
}

func NewAbstractModel(table string) *AbstractModel {
	return &AbstractModel{table: table}
}

func (m *AbstractModel) GetTableName() string {
	return m.table
}

func (m *AbstractModel) GetId() uuid.UUID {
	return m.ID
}

func (m *AbstractModel) SetId(id uuid.UUID) {
	m.ID = id
}
