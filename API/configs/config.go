package configs

import "os"

type Config struct{
	ServerPort string
}

func Loadconfig() Config {
	port := os.Getenv("PORT")
	if port == ""{
		port = ":8080"
	}
	return Config{
		ServerPort: port,
	}
}