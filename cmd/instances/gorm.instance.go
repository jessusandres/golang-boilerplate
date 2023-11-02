package instances

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm/logger"

	"lookerdevelopers/boilerplate/cmd/config"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	log.Println("🔨 Building pool...")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.EnvConfig.DBHost,
		config.EnvConfig.DBPort,
		config.EnvConfig.DBUsername,
		config.EnvConfig.DBPassword,
		config.EnvConfig.DBName,
		config.EnvConfig.DBSSLMode,
	)

	var db *gorm.DB
	var err error

	// Use the default driver when USE_SQL_CONNECTOR=false or not set
	if config.EnvConfig.UseSQLConnector {
		log.Println("☁️  Using cloudsqlpostgres driver")
		db, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: "cloudsqlpostgres",
			DSN:        dsn,
		}))

	} else {
		log.Println("⚙️ Using default postgres driver")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
		})

	}

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic("Failed to get db object")
	}

	sqlDB.SetMaxIdleConns(config.EnvConfig.DBMinConnections)

	sqlDB.SetMaxOpenConns(config.EnvConfig.DBMaxConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	log.Printf("DBLogger: %t", config.EnvConfig.DBLogger)

	if !config.EnvConfig.DBLogger {
		db.Logger = logger.Default.LogMode(logger.Silent)
	}

	DB = db
}
