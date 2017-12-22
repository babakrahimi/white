package app

type App struct {
	Repository *repository
}

func New() (*App, error) {
	r, err := newRepository()
	if err != nil {
		return nil, err
	}
	a := &App{r}
	return a, nil
}
