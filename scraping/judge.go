package scraping

import (
	"SAKU-PAY/database"
	"SAKU-PAY/model"
	"reflect"
)

func Compare_Notice_Judge(notice model.Notice) bool {
	db_notice_elements, _ := database.Get_Scrape_Notice()
	if len(db_notice_elements) == 0 {
		return false
	}

	db_notice := db_notice_elements[0]

	if reflect.DeepEqual(db_notice, notice) {
		return true
	}
	return false
}
