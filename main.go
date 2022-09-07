package main

import (
	"Belajar/CleanCode/config"
	"Belajar/CleanCode/factory"
	"Belajar/CleanCode/migration"
	"Belajar/CleanCode/utils/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	cfg := config.GetConfig()
	db := mysql.InitDBmySql(cfg)

	migration.InitMigrate(db)

	e := echo.New()
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))

}
