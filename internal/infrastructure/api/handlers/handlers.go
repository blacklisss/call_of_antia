package handlers

import (
	"antia/internal/usecases/app/repos/relationrepo"
	"antia/internal/usecases/app/repos/runerepo"
	"antia/internal/usecases/app/repos/teamrepo"
	"antia/internal/usecases/app/repos/userrepo"
)

type Handlers struct {
	us *userrepo.Users
	rs *runerepo.Runes
	ts *teamrepo.Teams
	rl *relationrepo.Relations
}

func NewHandlers(
	us *userrepo.Users,
	rs *runerepo.Runes,
	ts *teamrepo.Teams,
	rl *relationrepo.Relations,
) *Handlers {
	handlers := &Handlers{
		us,
		rs,
		ts,
		rl,
	}
	return handlers
}
