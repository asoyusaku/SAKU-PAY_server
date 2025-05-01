package variables

import (
	"gorm.io/gorm"
)

const Database_file string = "SAKU-PAY.db"
const Goods_list_url string = "https://store.plusmember.jp/sakurazaka46/"

var Database *gorm.DB
