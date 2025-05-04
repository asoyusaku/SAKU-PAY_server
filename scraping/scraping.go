package scraping

import (
	"SAKU-PAY/database"
	"SAKU-PAY/model"
	"SAKU-PAY/variables"
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

const WASTE_NUMBER int = 3 // waste_numberは表示順、表示件数、ALL ITEMSの3つをスキップするための変数

func Scrape_Goods() {
	var count int = 0
	var waste_count int = 0

	url := launcher.New().MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(variables.Goods_list_url)
	el := page.MustElementR("a", "ALL ITEMS")

	href, err := el.Attribute("href")
	if err != nil || href == nil {
		fmt.Println("hrefの取得に失敗しました")
		return
	}

	page = browser.MustPage(*href)

	page.MustWaitStable()

	title := page.MustInfo().Title
	fmt.Println("Page Title:", title)

	number_string := page.MustElement("p.item-num span").MustText()
	number, _ := strconv.Atoi(number_string)
	fmt.Println("Number of items:", number)
outer:
	for {
		elements_name := page.MustElements("p.tit")
		elements_price := page.MustElements("p.price span")
		figure := page.MustElements("figure.thumb img")

		for i := 0; i < len(elements_name); i++ {
			if count != number {
				goods_name := elements_name[i].MustText()
				price := elements_price[i].MustText()
				style := figure[i].MustAttribute("style")
				re := regexp.MustCompile(`url\((.*?)\)`)
				match := re.FindStringSubmatch(*style)
				if goods_name == "表示順" || goods_name == "表示件数" {
					if waste_count != WASTE_NUMBER {
						waste_count++
						continue
					}
					break
				}
				fmt.Println("Goods Name:", goods_name)
				fmt.Println("Price:", price)
				fmt.Println("Image URL:", match[1])
				goods := model.Goods{
					Name:  goods_name,
					Price: price,
					Image: match[1],
				}
				database.Add_Scrape_Goods(goods)
				count++
			} else {
				break outer
			}
		}
		page.MustElement("a.next").MustClick()
		page.MustWaitStable()
	}

}

func Scrape_Members() {
	url := launcher.New().MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(variables.Member_list_url)
	defer page.MustClose()

	title := page.MustInfo().Title
	fmt.Println("Page Title:", title)

	for {
		elements_name := page.MustElements("p.name")
		if len(elements_name) == 0 {
			continue
		} else {
			for i := 0; i < len(elements_name); i++ {
				member_name := elements_name[i].MustText()
				img := page.MustElement("img[alt='" + member_name + "']")
				member_photo := img.MustAttribute("src")
				fmt.Println("Member Name:", member_name)
				photo := fmt.Sprintf("%s%s", variables.Base_url, *member_photo)
				fmt.Println("Member Photo:", photo)
				member_photo = &photo
				member := model.Member{
					Name:    member_name,
					Picture: *member_photo,
				}
				database.Add_Scrape_Member(member)
			}
			break
		}
	}
	fmt.Println("")
	fmt.Println("")
	db, _ := database.GetMember()
	for _, member := range db {
		fmt.Println("Member Name:", member.Name)
		fmt.Println("Member Photo:", member.Picture)
	}

}

func Scrape_Goods_Notice() {
	url := launcher.New().MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(variables.Goods_list_url)
	defer page.MustClose()

	page.MustElement("div.area--news p.tit") //明示的に要素が表示されるまで待機

	elements_text := page.MustElements("div.area--news p.tit")
	elements_date := page.MustElements("div.area--news p.date")

	for count := 0; count < len(elements_text); count += 2 {
		text := elements_text[count].MustText()
		date := elements_date[count].MustText()
		notice := model.Notice{
			Text: text,
			Date: date,
		}
		if !Compare_Notice_Judge(notice) && (count == 0) {
			database.Add_Scrape_Notice(notice)
			Scrape_Goods()
			continue
		}
		database.Add_Scrape_Notice(notice)
	}
}
