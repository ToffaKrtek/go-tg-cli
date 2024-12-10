package environment

import (
	"github.com/joho/godotenv"
)

var envMap map[string]string

func Parse(envPath string) {
	envMap = make(map[string]string)
	if len(envPath) == 0 {
		envPath = "./.env"
	}

	envFile, err := godotenv.Read(envPath)
	if err == nil {
		envMap = envFile
	}
}

func Get(key string) string {
	if envMap == nil {
		Parse("")
	}
	if v, ok := envMap[key]; ok {
		return v
	}
	return ""
}
