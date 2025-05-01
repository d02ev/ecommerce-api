package db

import (
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"github.com/d02ev/ecommerce-api/pkg/config"
	"github.com/d02ev/ecommerce-api/pkg/logger"
)

var DB *gorm.DB;
type logrusWriter struct {}

func (lw *logrusWriter) Printf(format string, v ...interface{}) {
	logger.Log.Infof(format, v...);
}

func Init() {
	if err := config.Load(); err != nil {
		logger.Log.Fatalf("failed to load config: %v", err);
	}

	dsn := config.DBConnectionString();
	gormLog := gormLogger.New(
		&logrusWriter{},
		gormLogger.Config{
			SlowThreshold: time.Second,
			LogLevel:      gormLogger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful: false,
		},
	)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLog,
	})
	if err != nil {
		logger.Log.Fatalf("failed to connect to database: %v", err);
	}

	sqlDB, err := dbConn.DB();
	if err != nil {
		logger.Log.Fatalf("failed to get sqlDB: %v", err);
	}
	sqlDB.SetMaxIdleConns(10);
	sqlDB.SetMaxOpenConns(100);
	sqlDB.SetConnMaxLifetime(time.Hour);

	DB = dbConn;
	logger.Log.Info("database connection established");
}