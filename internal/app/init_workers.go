package app

// InitWorkers ...
func (t *App) initWorkers() []worker {
	return []worker{
		serveHTTP,
	}
}
