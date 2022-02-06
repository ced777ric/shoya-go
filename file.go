package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	BaseModel
	Url string
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = "file_" + uuid.New().String()
	return
}
