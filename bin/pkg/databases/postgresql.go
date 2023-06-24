package databases

import (
	"fmt"
	"time"

	"agree-agreepedia/bin/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgre() *gorm.DB {

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.GlobalEnv.PostgreSQL.Host,
		config.GlobalEnv.PostgreSQL.User,
		config.GlobalEnv.PostgreSQL.Password,
		config.GlobalEnv.PostgreSQL.DBName,
		config.GlobalEnv.PostgreSQL.Port,
		config.GlobalEnv.PostgreSQL.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database postgre")
	}

	postgresDb, err := db.DB()
	if err != nil {
		panic("Failed to create pool connection database postgre")
	}

	PostgresMaxLifeTime := time.Duration(config.GlobalEnv.PostgreSQL.MaxLifeTime)
	postgresDb.SetMaxOpenConns(config.GlobalEnv.PostgreSQL.MaxOpenConns)
	postgresDb.SetMaxIdleConns(config.GlobalEnv.PostgreSQL.MaxIdleConns)
	postgresDb.SetConnMaxLifetime(PostgresMaxLifeTime * time.Second)
	return db
}
