package scraping

import (
	"SAKU-PAY/variables"
	"fmt"
	"strconv"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

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
		elements := page.MustElements("p.tit")
		for _, element := range elements {
			if count != number {
				goods := element.MustText()
				if goods == "表示順" || goods == "表示件数" {
					if waste_count != 3 {
						waste_count++
						continue
					}
					break
				}
				fmt.Println("Goods Name:", goods)
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

	title := page.MustInfo().Title
	fmt.Println("Page Title:", title)
}
