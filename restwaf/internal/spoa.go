package internal

type application struct {
}

type SPOA struct {
	applications       map[string]*application
	defaultApplication string
}

func New() (*SPOA, error) {
	return nil, nil
}
