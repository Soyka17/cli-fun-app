package main

import (
	"ExperienceBank/internal/app"
	"ExperienceBank/internal/controller/logger"
	"ExperienceBank/internal/usecase/db"
	"ExperienceBank/internal/usecase/ui"
	"time"
)

func main() {
	logger := logger.NewLogger("info")
	db := db.NewDB(logger)
	ui := ui.NewCliUiRepositoryImpl(80, 70*time.Millisecond, logger)
	app := app.NewApp(ui, db, logger)

	app.Run()
}
