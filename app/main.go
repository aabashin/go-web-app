package main

import (
	"app/data"
	"app/handlers"

	"github.com/tsawler/celeritas"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
