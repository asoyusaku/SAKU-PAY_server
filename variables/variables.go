package variables

import (
	"gorm.io/gorm"
)

const Database_file string = "SAKU-PAY.db"
const Base_url string = "https://sakurazaka46.com"
const Goods_list_url string = "https://store.plusmember.jp/sakurazaka46/"
const Member_list_url string = "https://sakurazaka46.com/s/s46/search/artist"

var Database *gorm.DB
