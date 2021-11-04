package service

type Repository struct {
	*Tag
	*City
	*Event
	*EventPart
}

func NewRepositories(t TagRepository, c CityRepository, e EventRepository, ep EventPartRepository) *Repository {
	return &Repository{
		Tag:       newTagRepo(t),
		City:      newCity(c),
		Event:     newEvent(e),
		EventPart: newEventPart(ep),
	}
}
