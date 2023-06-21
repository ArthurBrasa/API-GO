package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  BDConfig
}

type APIConfig struct {
	Port string
}

type BDConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// função que sempre eh chamado no start das aplicações
func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

// Leitura do arquivo com os valores
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config) // Cria um ponteiro da struct
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = BDConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDB() BDConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
