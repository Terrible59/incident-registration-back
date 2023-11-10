package main

import (
	"database/sql"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	db "incidents_back/db/sqlc"
	"incidents_back/pkg/handlers"
	"incidents_back/pkg/utils"
	"time"
)

func main() {
	conn, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s dbname=%s user='%s' password=%s sslmode=disable",
			utils.GetEnv("DB_HOST", "localhost"),
			utils.GetEnv("DB_PORT", "5432"),
			utils.GetEnv("DB_NAME", "incident"),
			utils.GetEnv("DB_USER", "incident"),
			utils.GetEnv("DB_PASSWORD", "incident"),
		),
	)
	if err != nil {
		log.WithError(err).Fatal("failed to connect to DB")
	}

	repo := db.NewRepo(conn)

	app := fiber.New(fiber.Config{
		AppName:     "Incidents",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		IdleTimeout: time.Minute,
	})

	handlers.SetupRoutes(&handlers.SetupConfig{
		App:  app,
		Repo: repo,
	})

	if err := app.Listen(":8888"); err != nil {
		log.WithError(err).Fatal("server failed to start!")
	}
}
