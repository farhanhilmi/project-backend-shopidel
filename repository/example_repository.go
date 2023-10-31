package repository

import (
	"gorm.io/gorm"
)

type exampleRepository struct {
	db *gorm.DB
}
type ExampleRepository interface {
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {

}
