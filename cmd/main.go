package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/babulal107/go-cloud-native-app/internal/config"
	"github.com/babulal107/go-cloud-native-app/internal/migration"
	"github.com/babulal107/go-cloud-native-app/internal/router"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var dbInstance *sql.DB

// In Go, sync.Once is a thread-safe way to ensure that initialization happens only once, even in concurrent environments.
var dbOnce sync.Once

func main() {

	// Initialize the database connection
	dbInstance = GetDBInstance()

	// Run migration to create users table if not exist
	migration.CreateUserTable(dbInstance)

	// AppContainer
	appContainer := config.AppContainer{
		DB: dbInstance,
	}

	// Init router
	ginRouter := router.Init(appContainer)

	// Define the server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: ginRouter,
	}

	// Run the server in a goroutine so it doesn’t block
	go func() {
		log.Println("Server started on :8080")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown
	GraceFulShutDown(srv)

}

// GetDBInstance provides a singleton database connection pool
// The GetDBInstance function ensures that only one database connection pool is created, using sync.Once.
func GetDBInstance() *sql.DB {
	dbOnce.Do(func() {
		var err error

		dbHost := "postgres_db" // Use the service name
		dbPort := "5432"
		dbUser := "root"        // Explicitly provide username
		dbPassword := "user123" // Explicitly provide password
		dbName := "users_db"

		// Correct connection string
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName,
		)

		// Initialize the database connection
		dbInstance, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}

		// Optionally ping the database to check connection
		if err = dbInstance.Ping(); err != nil {
			log.Fatal("Database is unreachable:", err)
		}
		fmt.Println("Database connection established.")
	})
	return dbInstance
}

func GraceFulShutDown(srv *http.Server) {

	// Graceful shutdown setup
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // Catch system signals

	<-quit // Wait for a shutdown signal
	log.Println("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")

}
