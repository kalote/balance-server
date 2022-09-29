package conf

type Config struct {
	Port      int    `envconfig:"PORT" required:"true" default:"3000"`
	InfuraKey string `envconfig:"APIKEY" required:"true"`
}
