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

func AddUser(response model.Response) error {
	user := model.User{
		ID:      response.Sub,
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

func AddOshimen(sub string, oshimen string) error {
	var user model.User
	if err := variables.Database.Where("id = ?", sub).First(&user).Error; err != nil {
		return err
	}
	member, err := GetMember(oshimen)
	if err != nil {
		return err
	}
	picture := member.Picture

	oshimenList := model.Oshimen{
		Name:    oshimen,
		Picture: picture,
	}

	user, err = GetUser(sub)
	if err != nil {
		return err
	}

	if err := variables.Database.Model(&user).Association("Oshimen").Append(&oshimenList); err != nil {
		return err
	}
	if err := variables.Database.Save(&user).Error; err != nil {
		return err
	}

	return nil

}

func AddMember(member model.Member) error {
	if err := variables.Database.Create(&member).Error; err != nil {
		return err
	}
	return nil
}

func GetMember(name string) (model.Member, error) {
	var member model.Member
	if err := variables.Database.Where("name = ?", name).First(&member).Error; err != nil {
		return model.Member{}, err
	}
	return member, nil
}

func GetOshimen(sub string) ([]model.Oshimen, error) {
	var user model.User

	if err := variables.Database.Where("id = ?", sub).First(&user).Error; err != nil {
		return nil, err
	}
	var oshimenList []model.Oshimen
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

func Get_Scrape_Member() ([]model.Member, error) {
	var members []model.Member
	if err := variables.Database.Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil

}

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
	database.AutoMigrate(&model.Member{})

	fmt.Println("Database connected successfully")

}
