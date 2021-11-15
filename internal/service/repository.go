package service

type Repository struct {
	*Tag
	*Event
	*EventPart
}

func NewServices(t TagRepository, e EventRepository, ep EventPartRepository) *Repository {
	return &Repository{
		Tag:       newTagRepo(t),
		Event:     newEvent(e),
		EventPart: newEventPart(ep),
	}
}
