package migrate

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/watiko/bibliotheca-go/internal/pkg"
)

type Env struct {
	pkg.DBEnv
}

func main() {
	var env Env

	err := envconfig.Process("", &env)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	fmt.Printf("host: %s, name: %s", env.DbHost, env.DbName)
}
