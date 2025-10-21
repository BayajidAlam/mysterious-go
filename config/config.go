package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	HOST          string
	PORT          int
	NAME          string
	USER          string
	PASSWORD      string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(
			"Failed to load the env variables: ", err,
		)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required!")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required!")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http port name is required!")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println(
			"Port must be number: ", err,
		)
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Secret key is required!")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("Host is required!")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("Db port is required!")
		os.Exit(1)
	}

	dbPrt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println(
			"Port must be number: ", err,
		)
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("NAME is required!")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("USER is required!")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("PASSWORD is required!")
		os.Exit(1)
	}

	enableSsLMode := os.Getenv("EnableSSLMODE")
	enableSSLMode, err := strconv.ParseBool(enableSsLMode)
	if err != nil {
		fmt.Println("Invalid enable ssl mode value", err)
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		HOST:          dbHost,
		PORT:          int(dbPrt),
		NAME:          dbName,
		USER:          dbUser,
		PASSWORD:      dbPassword,
		EnableSSLMODE: enableSSLMode,
	}
	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
