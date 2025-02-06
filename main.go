package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/techerfan/2DCH7-20059/delivery/httpserver/fiber"
	"github.com/techerfan/2DCH7-20059/pkg/logger"
	"github.com/techerfan/2DCH7-20059/pkg/myjwt"
	"github.com/techerfan/2DCH7-20059/repository/postgres"
	"github.com/techerfan/2DCH7-20059/repository/redis"
	"github.com/techerfan/2DCH7-20059/service/userservice"
)

// A day is consisted of 86400 seconds
const TokenExpirationTime = 86400

func main() {
	// Read environment variables
	redisPortStr := os.Getenv("REDIS_PORT")
	postgresPortStr := os.Getenv("POSTGRES_PORT")
	portStr := os.Getenv("PORT")
	dbName := os.Getenv("DBName")
	dbUser := os.Getenv("DBUser")
	dbPassFile := os.Getenv("POSTGRES_PASSWORD_FILE")
	jwtSecretFile := os.Getenv("JWT_SECRET_FILE")

	// Parse ports
	redisPort, _ := strconv.ParseInt(redisPortStr, 10, 64)
	postgresPort, _ := strconv.ParseInt(postgresPortStr, 10, 64)
	port, _ := strconv.ParseInt(portStr, 10, 64)

	// Read secrets
	jwtSecret, _ := os.ReadFile(jwtSecretFile)
	postgresPass, _ := os.ReadFile(dbPassFile)

	// Make an instance of JWT
	tokenGenerator := myjwt.New()
	tokenGenerator.SetSecret(jwtSecret)
	tokenGenerator.SetClaims("userId", "exp")

	// Make an instance of the postgres database
	postgresDB := postgres.New(postgres.Config{
		Username: dbUser,
		Password: string(postgresPass),
		DBName:   dbName,
		Host:     "postgres",
		Port:     int(postgresPort),
	})

	// Make an instance of the redis cache store
	redisDB := redis.New(redis.Config{
		Host: "redis",
		Port: int(redisPort),
	})

	// User service instance
	userService := userservice.New(TokenExpirationTime, tokenGenerator, postgresDB, redisDB)

	// TODO: this must be replaced with a real logger
	var dummyLogger logger.Logger = logger.DummyLogger{}

	// HTTP server instance
	httpServer := fiber.New(tokenGenerator, dummyLogger, TokenExpirationTime, userService)

	// Start the server
	if err := httpServer.Start(fmt.Sprintf("%d", port)); err != nil {
		panic(err)
	}
}
