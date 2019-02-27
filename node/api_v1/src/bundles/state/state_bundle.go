package state

import (
	"net/http"

	"../../core"
)

// StateBundle handle state of service
type StateBundle struct {
	routes []core.Route
}

// NewStateBundle instance
func NewStateBundle() core.Bundle {

	routes := []core.Route{
		core.Route{
			Method:  http.MethodGet,
			Path:    "/ping",
			Handler: ping,
		},
	}

	return &StateBundle{
		routes: routes,
	}
}

// GetRoutes implement interface core.Bundle
func (b *StateBundle) GetRoutes() []core.Route {
	return b.routes
}
