package database

import (
	"SAKU-PAY/model"
	"SAKU-PAY/variables"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func AddIdToken(token model.IdToken) error { //complete
	idtoken := model.IdToken{
		IdToken: token.IdToken,
	}
	if err := variables.Database.Create(&idtoken).Error; err != nil {
		return err
	}
	return nil
}

func GetIdToken() (model.IdToken, error) {
	var idtoken model.IdToken
	if err := variables.Database.First(&idtoken).Error; err != nil {
		return model.IdToken{}, err
	}
	return idtoken, nil
}

func AddUser(response model.Response) error { //complete
	user := model.User{
		ID:      "12345",
		Name:    response.Name,
		Picture: response.Picture,
		Email:   response.Email,
	}

	if err := variables.Database.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func GetUser(sub string) (model.User, error) {
	var user model.User
	if err := variables.Database.Where("id = ?", sub).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

func AddOshimen(sub string, oshimen model.Member) error { //complete
	var user model.User
	var existingMember model.Member // Check if the member already exists

	if err := variables.Database.Where("id = ?", sub).First(&user).Error; err != nil {
		return err
	}
	if err := variables.Database.Where("name = ?", oshimen.Name).First(&existingMember).Error; err == nil {
		return nil // Member already exists, no need to create a new record
	}
	if err := variables.Database.Model(&user).Association("Oshimen").Append(&oshimen); err != nil {
		return err
	}
	if err := variables.Database.Save(&user).Error; err != nil {
		return err
	}
	return nil

}

func AddMember(member model.Member) error { //complete
	if err := variables.Database.Create(&member).Error; err != nil {
		return err
	}
	return nil
}

func GetMember() ([]model.Member, error) { //complete
	var members []model.Member
	if err := variables.Database.Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func GetOshimen(sub string) ([]model.Member, error) { //complete
	var user model.User

	if err := variables.Database.Where("id = ?", sub).First(&user).Error; err != nil {
		return nil, err
	}

	var oshimenList []model.Member
	if err := variables.Database.Model(&user).Association("Oshimen").Find(&oshimenList); err != nil {
		return nil, err
	}

	return oshimenList, nil
}

func AddGoods(sub string, goods model.Goods) error {
	var user model.User

	//userのgoodslistにgoodsを追加する
	if err := variables.Database.Where("id = ?", sub).First(&user).Error; err != nil {
		return err
	}

	if err := variables.Database.Model(&user).Association("Goods").Append(&goods); err != nil {
		return err
	}

	if err := variables.Database.Save(&user).Error; err != nil {
		return err
	}

	return nil

}

func GetGoods(sub string) ([]model.Goods, error) {
	var user model.User

	if err := variables.Database.Where("id = ?", sub).First(&user).Error; err != nil {
		return nil, err
	}

	var goodsList []model.Goods
	if err := variables.Database.Model(&user).Association("Goods").Find(&goodsList); err != nil {
		return nil, err
	}

	return goodsList, nil
}

func Add_Scrape_Goods(goods model.Goods) error {
	if err := variables.Database.Create(&goods).Error; err != nil {
		return err
	}
	return nil
}

func Add_Scrape_Member(member model.Member) error {
	// Check if the member already exists
	var existingMember model.Member
	if err := variables.Database.Where("name = ?", member.Name).First(&existingMember).Error; err == nil {
		return nil // Member already exists, no need to create a new record
	}

	// If the member does not exist, create a new record
	if err := variables.Database.Create(&member).Error; err != nil {
		return err
	}

	return nil
}

func Database() {
	database, err := gorm.Open(sqlite.Open(variables.Database_file), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	variables.Database = database

	database.AutoMigrate(&model.IdToken{})
	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Member{})
	database.AutoMigrate(&model.Goods{})

	fmt.Println("Database connected successfully")

}
