package app

type Handler struct {
}

func New() (*Handler, error) {
	return &Handler{}, nil
}
