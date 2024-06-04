package middleware

import (
	"app/data"

	"github.com/tsawler/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
