package main

import (
	"os"

	"github.com/jgourdin/42shop_api/model"
)

func main() {
	a := model.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	a.Run(":8001")
}
