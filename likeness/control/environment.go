package control

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port      uint16 `envconfig:"PORT" default:"3000"`
	Host      string `envconfig:"HOST" required:"true`
	AdminPort uint16 `envconfig:"ADMIN_PORT" default:"3001`
}

func main() {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	log.Println(c)
}
