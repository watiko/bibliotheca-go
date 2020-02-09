package pkg

type DBEnv struct {
	DbHost string `required:"true" split_words:"true"`
	DbPort string `required:"true" split_words:"true"`
	DbUser string `required:"true" split_words:"true"`
	DbPass string `required:"true" split_words:"true"`
	DbName string `required:"true" split_words:"true"`
}
