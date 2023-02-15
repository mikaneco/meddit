package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	// Server
	ServerPort string

	// Database
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string

	// JWT
	JWTSecretKey string
	JWTDuration  time.Duration

	// Stripe
	StripeSecretKey string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// Server
	ServerPort = os.Getenv("SERVER_PORT")

	// Database
	DatabaseHost = os.Getenv("DATABASE_HOST")
	DatabasePort = os.Getenv("DATABASE_PORT")
	DatabaseUser = os.Getenv("DATABASE_USER")
	DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	DatabaseName = os.Getenv("DATABASE_NAME")

	// JWT
	JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	JWTDuration, err = time.ParseDuration(os.Getenv("JWT_DURATION"))
	if err != nil {
		panic(err)
	}
	// Stripe
	StripeSecretKey = os.Getenv("STRIPE_SECRET_KEY")
}
