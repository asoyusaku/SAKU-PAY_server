package main

import (
	"SAKU-PAY/api"
	"SAKU-PAY/database"
	"SAKU-PAY/scraping"
)

func main() {
	database.Database()
	api.Api()
	// scraping.Scrape_Members()
	// scraping.Scrape_Goods()
	scraping.Scrape_Members()
}
