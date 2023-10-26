package model

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Task struct {
	ID       uuid.UUID
	Name     string
	Finished bool
}
