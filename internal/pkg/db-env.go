package pkg

type DBEnv struct {
	DbURL string `required:"true" split_words:"true"`
}
