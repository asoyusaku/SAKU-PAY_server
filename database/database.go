package database

import (
	"SAKU-PAY/model"
	"SAKU-PAY/response"
	"SAKU-PAY/variables"
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

func AddUser(token model.IdToken) error { //complete
	response, err := response.LineVerify(token.IdToken)
	if err != nil {
		return err
	}

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

func AddOshimen(sub string, oshimen model.Member) error {
	var user model.User
	var existingMember model.Member

	if err := variables.Database.First(&user, "id = ?", sub).Error; err != nil {
		return err
	}

	if err := variables.Database.Where("name = ?", oshimen.Name).First(&existingMember).Error; err != nil {
		if err := variables.Database.Create(&oshimen).Error; err != nil {
			return err
		}
		existingMember = oshimen
	}

	if err := variables.Database.Model(&user).Association("Oshimen").Append(&existingMember); err != nil {
		return err
	}
	return nil
}

func GetOshimen(sub string) ([]model.Member, error) { //complete
	var user model.User
	var oshimen_list []model.Member

	if err := variables.Database.Preload("Oshimen").First(&user, "id = ?", sub).Error; err != nil {
		return nil, err
	}
	if err := variables.Database.Model(&user).Association("Oshimen").Find(&oshimen_list); err != nil {
		return nil, err
	}
	return oshimen_list, nil

}

func DeleteOshimen(sub string, oshimen model.Member) error { //complete
	var user model.User

	if err := variables.Database.Where("id = ?", sub).First(&user).Error; err != nil {
		return err
	}

	if err := variables.Database.Model(&user).Association("Oshimen").Delete(&oshimen); err != nil {
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

func GetGoods(sub string) ([]model.Response_Goods, error) { //complete
	var purchases []model.Purchase
	var goods []model.Response_Goods

	// ユーザーが購入したすべての購入データを取得
	if err := variables.Database.Where("user_id = ?", sub).Find(&purchases).Error; err != nil {
		return nil, err
	}

	fmt.Println("Purchases:", purchases)
	// 各購入データに関連するグッズを取得
	for _, purchase := range purchases {
		var good model.Goods
		if err := variables.Database.Where("name = ?", purchase.GoodsName).First(&good).Error; err != nil {
			return nil, err
		}
		goods = append(goods, model.Response_Goods{
			Goods:    good,
			Quantity: purchase.Quantity,
		})
	}

	return goods, nil
}

func UpdateGoods(request model.Request_Purchase) error {
	var user model.User
	var goods model.Goods
	var purchase model.Purchase

	if err := variables.Database.Where("user_id = ? AND goods_name = ?", request.UserId, request.GoodsName).First(&purchase).Error; err != nil {
		purchase := model.Purchase{
			UserId:    user.ID,
			GoodsName: goods.Name,
			Quantity:  request.Quantity,
		}

		if err := variables.Database.Create(&purchase).Error; err != nil {
			fmt.Println("Failed to create purchase:", err)
			return err
		}
		return nil
	} else {
		purchase.Quantity += request.Quantity
		if err := variables.Database.Save(&purchase).Error; err != nil {
			return err
		}
		return nil
	}
}

// データベース上にある全てのグッズを取得する
func GetAllGoods() ([]model.Goods, error) {
	var goods []model.Goods
	if err := variables.Database.Find(&goods).Error; err != nil {
		return nil, err
	}
	return goods, nil
}

func Add_Scrape_Goods(goods model.Goods) error { //complete
	if strings.Contains(goods.Name, "【会場受取】") || strings.Contains(goods.Name, "【通常配送】") {
		goods.Name = strings.ReplaceAll(goods.Name, "【会場受取】", "")
		goods.Name = strings.ReplaceAll(goods.Name, "【通常配送】", "")
	}
	if err := variables.Database.Create(&goods).Error; err != nil {
		return err
	}
	return nil
}

func Add_Scrape_Member(member model.Member) error { //complete
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

func Add_Scrape_Notice(notice model.Notice) error { //complete
	var existingNotice model.Notice
	if err := variables.Database.Where("text = ?", notice.Text).First(&existingNotice).Error; err == nil {
		return nil
	}
	if err := variables.Database.Create(&notice).Error; err != nil {
		return err
	}
	fmt.Println("success add scrape notice")
	return nil
}

func Get_Scrape_Notice() ([]model.Notice, error) { //complete
	var notices []model.Notice
	if err := variables.Database.Find(&notices).Error; err != nil {
		return nil, err
	}
	fmt.Println("success get scrape notice")
	return notices, nil
}

func Get_TotalCost(sub string) (int, error) { //complete
	var total_cost int
	var purchases []model.Purchase

	if err := variables.Database.Where("user_id = ?", sub).Find(&purchases).Error; err != nil {
		return 0, err
	}

	fmt.Println("Purchases:", purchases)
	for _, purchase := range purchases {
		fmt.Println("Purchase:", purchase)
		var goods model.Goods
		if err := variables.Database.Where("name = ?", purchase.GoodsName).First(&goods).Error; err != nil {
			fmt.Println("Failed to find goods:", err)
			return 0, err
		}
		price, _ := ExtractPrice(goods.Price)
		total_cost += purchase.Quantity * price
		fmt.Println("Goods name:", goods.Name)
		fmt.Println("Goods price:", price)
		fmt.Println("Quantity:", purchase.Quantity)
	}
	fmt.Println("Total cost:", total_cost)
	return total_cost, nil
}

func ExtractPrice(priceStr string) (int, error) { //complete
	// 正規表現で数値部分を抽出
	re := regexp.MustCompile(`[0-9,]+`)
	matches := re.FindString(priceStr)
	if matches == "" {
		return 0, fmt.Errorf("no valid price found in string")
	}

	// カンマを削除して数値に変換
	cleaned := regexp.MustCompile(`,`).ReplaceAllString(matches, "")
	price, err := strconv.Atoi(cleaned)
	if err != nil {
		return 0, err
	}

	return price, nil
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
	database.AutoMigrate(&model.Notice{})
	database.AutoMigrate(&model.Purchase{})

	fmt.Println("Database connected successfully")

}
