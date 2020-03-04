package types

type AppContext struct {
	Debug  bool
	Commit string
}

func NewAppContext(env string, commit string) *AppContext {
	return &AppContext{
		Debug:  env == "dev",
		Commit: commit,
	}
}
