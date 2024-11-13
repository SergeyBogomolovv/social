package app

type App struct {
	addr string
}

func NewApp(addr string) *App {
	return &App{
		addr: addr,
	}
}
