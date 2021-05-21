package main

import "os"

func main() {
	api := API{}
	api.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	api.Run(":8010")
}
