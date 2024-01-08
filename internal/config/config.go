package config

import "github.com/joho/godotenv"

func LoadEnvConfig(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
