package database

import (
	"SAKU-PAY/model"
	"SAKU-PAY/variables"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AddIdToken(token model.IdToken) error {
	idtoken := model.IdToken{
		IdToken: token.IdToken,
	}

	if err := variables.Database.Create(&idtoken).Error; err != nil {
		return err
	}

	return nil

}

// func GetIdToken() (model.IdToken, error) {
// }

// func AddUser(user model.User) error {
// }

// func GetUser() (model.User, error) {
// }

// func AddOshimen(oshimen model.Oshimen) error {
// }

// func GetOshimen() ([]model.Oshimen, error) {
// }

// func AddGoods(goods model.Goods) error {
// }

// func GetGoods() ([]model.Goods, error) {
// }

func Database() {
	database, err := gorm.Open(sqlite.Open(variables.Database_file), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	variables.Database = database

	database.AutoMigrate(&model.IdToken{})
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Oshimen{})
	database.AutoMigrate(&model.Goods{})

	fmt.Println("Database connected successfully")

}
