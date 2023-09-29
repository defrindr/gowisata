// config/config.go

package config

type AppConfig struct {
    DBUser     string
    DBPassword string
    DBName     string
    DBHost     string
    DBPort     string
    SecretKey  string
}

var Config = AppConfig{
    DBUser:     "your_database_user",
    DBPassword: "your_database_password",
    DBName:     "your_database_name",
    DBHost:     "maria_app",
    DBPort:     "3306",
    SecretKey:  "your_secret_key",
}
