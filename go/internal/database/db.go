package database

import (
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func db(serviceURI string) *gorm.DB {
	
	conn, err := url.Parse(serviceURI)
	if err != nil {
		log.Fatalf("failed to parse service URI: %v", err)
	}

	// Ensure SSL connection using the provided SSL root certificate
	conn.RawQuery = "sslmode=verify-ca&sslrootcert=cert/ca.pem"

	dsn := conn.String()

	// // Load SSL root certificate from file
	// sslRootCert := os.Getenv("SSL_ROOT_CERT")
	// if sslRootCert == "" {
	// 	sslRootCert = "cert/ca.pem" // Ensure this file is correctly referenced
	// }

	// dsn += fmt.Sprintf("&sslrootcert=%s", sslRootCert)

	// Initialize GORM with the PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Test the connection by querying the version
	var result string
	row := db.Raw("SELECT version()").Row()
	if err := row.Scan(&result); err != nil {
		log.Fatalf("failed to query database version: %v", err)
	}
	fmt.Printf("Version: %s\n", result)
	return db
}
