package config

import (
	"agree-agreepedia/bin/config/key"
	"crypto/rsa"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	RootApp    string
	HTTPPort   uint16
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey

	BasicAuth struct {
		Username      string
		Password      string
		AdminUsername string
		AdminPassword string
	}
	PostgreSQL struct {
		Host         string
		User         string
		Password     string
		DBName       string
		Port         uint16
		SSLMode      string
		MaxIdleConns int
		MaxOpenConns int
		MaxLifeTime  int
	}
	EncodeDecodeKey struct {
		EncodeKey string
		IVKey     string
	}
}

// GlobalEnv global environment
var GlobalEnv Env

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	rootApp := strings.TrimSuffix(path, "/bin/config")
	os.Setenv("APP_PATH", rootApp)
	GlobalEnv.RootApp = rootApp
	GlobalEnv.PrivateKey = key.LoadPrivateKey()
	GlobalEnv.PublicKey = key.LoadPublicKey()

	loadGeneral()
	loadBasicAuth()
	loadPostgreSQL()
	loadEncodeDecodeKey()
}

func loadGeneral() {
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	GlobalEnv.HTTPPort = uint16(port)
}

func loadBasicAuth() {
	GlobalEnv.BasicAuth.Username = os.Getenv("BASIC_AUTH_USERNAME")
	GlobalEnv.BasicAuth.Password = os.Getenv("BASIC_AUTH_PASSWORD")
}

func loadPostgreSQL() {
	GlobalEnv.PostgreSQL.Host = os.Getenv("POSTGRES_DB_HOST")
	GlobalEnv.PostgreSQL.User = os.Getenv("POSTGRES_DB_USER")
	GlobalEnv.PostgreSQL.Password = os.Getenv("POSTGRES_DB_PASSWORD")
	GlobalEnv.PostgreSQL.DBName = os.Getenv("POSTGRES_DB_NAME")
	Portpostgre, _ := strconv.Atoi(os.Getenv("POSTGRES_DB_PORT"))
	GlobalEnv.PostgreSQL.Port = uint16(Portpostgre)
	GlobalEnv.PostgreSQL.SSLMode = os.Getenv("POSTGRES_DB_SSLMODE")
}

func loadEncodeDecodeKey() {
	GlobalEnv.EncodeDecodeKey.EncodeKey = os.Getenv("ENCODE_KEY")
	GlobalEnv.EncodeDecodeKey.IVKey = os.Getenv("IV_KEY")
}
