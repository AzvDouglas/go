package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// Config is the configuration for the API
type Config struct {
	DBDriver string 	`mapstructure:"DB_DRIVER"`
	DBHost   string 	`mapstructure:"DB_HOST"`
	DBPort   string 	`mapstructure:"DB_PORT"`
	DBUser   string 	`mapstructure:"DB_USER"`
	DBName   string 	`mapstructure:"DB_NAME"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret string 	`mapstructure:"JWT_SECRET"`
	JwtExpirationTime int `mapstructure:"JWT_EXPIRATION_TIME"`
	TokenAuth	*jwtauth.JWTAuth
}
	
func init() {		// init() is called before main()
	LoadConfig(".")
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(path + "env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
