package config

type Config struct {
	Server      Server
	Database    Database
	Redis       Redis
	PKI         PKI
	GCSConfig   GCSConfig
	MinioConfig MinioConfig
}

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Redis struct {
	Host string
	Port string
	Db   string
}

type PKI struct {
	PrivateKey string
	PublicKey  string
}

type GCSConfig struct {
	UseGCS         bool
	CredentialFile string
}

type MinioConfig struct {
	UseMinio        bool
	Endpoint        string
	AccessKey       string
	SecretAccessKey string
	SSL             bool
}
