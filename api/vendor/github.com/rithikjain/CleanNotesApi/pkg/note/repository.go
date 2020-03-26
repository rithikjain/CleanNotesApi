package note

type Repository interface {
	CreateNote(note *Note) (*Note, error)

	GetAllNotes(userID float64) (*[]Note, error)
}
