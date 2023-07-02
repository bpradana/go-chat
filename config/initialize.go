package config

import (
	"os"
)

func Initialize() Config {
	cfg := Config{
		Server: Server{
			Port: os.Getenv("PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Name:     os.Getenv("DB_NAME"),
		},
		Redis: Redis{
			Host: os.Getenv("REDIS_HOST"),
			Port: os.Getenv("REDIS_PORT"),
			Db:   os.Getenv("REDIS_DB"),
		},
		PKI: PKI{
			PrivateKey: os.Getenv("PRIVATE_KEY"),
			PublicKey:  os.Getenv("PUBLIC_KEY"),
		},
		GCSConfig: GCSConfig{
			UseGCS:         os.Getenv("USE_GCS") == "true",
			CredentialFile: os.Getenv("GCS_CREDENTIAL_FILE"),
		},
		MinioConfig: MinioConfig{
			UseMinio:        os.Getenv("USE_MINIO") == "true",
			Endpoint:        os.Getenv("MINIO_ENDPOINT"),
			AccessKey:       os.Getenv("MINIO_ACCESS_KEY"),
			SecretAccessKey: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
			SSL:             os.Getenv("MINIO_SSL") == "true",
		},
	}

	return cfg
}
