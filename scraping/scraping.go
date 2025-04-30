package scraping

import (
	"SAKU-PAY/variables"
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Scrape() {

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

	for {
		elements := page.MustElements("p.tit")
		for _, element := range elements {
			goods := element.MustText()
			if goods == "表示順" || goods == "表示件数" {
				continue
			}
			fmt.Println("Goods Name:", goods)
		}
		page.MustElement("a.next").MustClick()
		page.MustWaitStable()
	}

}

/*
	//櫻坂46のグッズの一覧を取得するurl
	curl -i "https://store.plusmember.jp/sakurazaka46/?srsltid=AfmBOopentiCrs9JWLjSsGbqTglgZFS3R3aryUktJFzNOPof_SbKKSnv"

	//グッズの画像を取得するurl(例：向井純葉のペンライト)
	url(https://storage-store.plusmember.jp/upload/images/sakura_PenlightVer2_Zq7CK6Sw_Mukai1.jpg)
*/
