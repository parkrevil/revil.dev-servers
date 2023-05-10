package main

type App struct {
	config Config
	server Server
}

func NewApp(config Config, server Server) App {
	config.print()

	return App{
		config: config,
		server: server,
	}
}

func (a *App) start() error {
	if err := a.server.start(); err != nil {
		return err
	}

	return nil
}

func (a *App) shutdown() error {
	if err := a.server.shutdown(); err != nil {
		return err
	}

	return nil
}
