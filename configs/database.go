package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}
	sqlDB.Close()
}

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		logLevelStr = strconv.Itoa(int(logger.Info))
	}
	logLevel, err := strconv.Atoi(logLevelStr)
	if err != nil {
		panic("Failed to load log level")
	}

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.LogLevel(logLevel),
			Colorful:      true,
		},
	)

	var (
		createBatchSize = os.Getenv("CREATE_BATCH_SIZE")
		maxOpenConns    = os.Getenv("MAX_OPEN_CONNS")
		maxIdleConns    = os.Getenv("MAX_IDLE_CONNS")
		connMaxLifetime = os.Getenv("CONN_MAX_LIFETIME")
		connMaxIdleTime = os.Getenv("CONN_MAX_IDLE_TIME")
	)
	if createBatchSize == "" {
		createBatchSize = "5000"
	}
	if maxOpenConns == "" {
		maxOpenConns = "100"
	}
	if maxIdleConns == "" {
		maxIdleConns = "50"
	}
	if connMaxLifetime == "" {
		connMaxLifetimeDur := 24 * time.Hour
		connMaxLifetime = connMaxLifetimeDur.String()
	}
	if connMaxIdleTime == "" {
		connMaxIdleTimeDur := time.Hour
		connMaxIdleTime = connMaxIdleTimeDur.String()
	}

	createBatchSizeInt, err := strconv.Atoi(createBatchSize)
	if err != nil {
		panic("Failed to load create batch size")
	}
	maxOpenConnsInt, err := strconv.Atoi(maxOpenConns)
	if err != nil {
		panic("Failed to load max open conns")
	}
	maxIdleConnsInt, err := strconv.Atoi(maxIdleConns)
	if err != nil {
		panic("Failed to load max idle conns")
	}
	connMaxLifetimeDur, err := time.ParseDuration(connMaxLifetime)
	if err != nil {
		panic("Failed to load conn max lifetime")
	}
	connMaxIdleTimeDur, err := time.ParseDuration(connMaxIdleTime)
	if err != nil {
		panic("Failed to load conn max idle time")
	}

	gormConfig := &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 dbLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		CreateBatchSize: createBatchSizeInt,
	}
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		dbHost = os.Getenv("DB_HOST_DOCKER")
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser,
			dbPass,
			dbHost,
			dbPort,
			dbName,
		)

		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
		if err != nil {
			panic("Failed to create a connection to database")
		}
	}

	dbCli, err := db.DB()
	if err != nil {
		panic("Cannot init DB")
	}

	dbCli.SetMaxOpenConns(maxOpenConnsInt)
	dbCli.SetMaxIdleConns(maxIdleConnsInt)
	dbCli.SetConnMaxLifetime(connMaxLifetimeDur)
	dbCli.SetConnMaxIdleTime(connMaxIdleTimeDur)

	return db

}
