package note

type Service interface {
	CreateNote(note *Note) (*Note, error)

	ShowAllNotes(userID float64) (*[]Note, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s service) CreateNote(note *Note) (*Note, error) {
	return s.repo.CreateNote(note)
}

func (s service) ShowAllNotes(userID float64) (*[]Note, error) {
	return s.repo.GetAllNotes(userID)
}

func (s *service) GetRepo() Repository {
	return s.repo
}
