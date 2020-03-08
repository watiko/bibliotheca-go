package types

import "context"

type AppContext struct {
	Debug  bool
	Commit string
	context.Context
}

// TODO: logger
func NewAppContext(env string, commit string) *AppContext {
	return &AppContext{
		Debug:  env == "dev",
		Commit: commit,
	}
}
