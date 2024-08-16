package config

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		PublicHost: "http://localhost",
		Port:       ":9002",
		DBUser:     "krish",
		DBPassword: "krish",
		DBAddress:  "127.0.0.1:3306",
		DBName:     "blog_website",
	}
}
