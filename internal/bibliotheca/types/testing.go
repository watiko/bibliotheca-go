package types

func NewTestAppContext() *AppContext {
	return &AppContext{
		Debug:  true,
		Commit: "testing",
	}
}
