package config

import (
	"os"
)

var (
	//ENV    string
	URL    string
	ID     string
	Secret string
)

func GetEnvDefault(key string, defVal string) string {

	val, ex := os.LookupEnv(key)
	if !ex {
		val = defVal
	}
	return val

}

func SetEnvionment() {
	//ENV = GetEnvDefault("ENV", "local")

	URL = GetEnvDefault("URL", "http://localhost:8081/callback")
	ID = GetEnvDefault("ID", "")
	Secret = GetEnvDefault("Secret", "")

}
