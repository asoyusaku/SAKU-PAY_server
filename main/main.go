package main

import (
	api "SAKU-PAY/api"
	database "SAKU-PAY/db"
)

func main() {
	database.Database()
	api.Api()
	// scraping.Scrape()
}
