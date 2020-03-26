package note

import (
	"github.com/jinzhu/gorm"
	"github.com/rithikjain/CleanNotesApi/pkg"
)

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateNote(note *Note) (*Note, error) {
	err := r.DB.Create(note).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return note, nil
}

func (r *repo) GetAllNotes(userID float64) (*[]Note, error) {
	var notes []Note
	err := r.DB.Where("user_id = ?", userID).Find(&notes).Error
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return &notes, nil
}
