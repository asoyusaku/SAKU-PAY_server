package main

import (
	"SAKU-PAY/api"
	"SAKU-PAY/database"
	"SAKU-PAY/scraping"
)

func main() {
	database.Database()
	scraping.Scrape_Goods_Notice()
	api.Api()
}
