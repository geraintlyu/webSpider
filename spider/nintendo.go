package spider

import (
	"Spider/Spider/src/modul"
	"fmt"
	"regexp"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly/v2"
	"xorm.io/xorm"
)

// insert to database
func saveDatabase(nintendo *modul.Nintendo, i *int) {
	engine, err := xorm.NewEngine("mysql", "root:password@tcp(ip:port)/databasename")
	if err != nil {
		fmt.Printf("Failed to dataBase engine: %v\n", err)
	}
	_, err2 := engine.Insert(*nintendo)
	if err2 != nil {
		fmt.Printf("Fail to insert data %v th", *i)
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("the %vth insert successfully\n", *i)
	}
	*i++
}

func Nintendo(wg *sync.WaitGroup) {
	defer wg.Done()
	i := 1
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("user-agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39`)
	})

	c.OnHTML("div[class='category-product-item-info']", func(h *colly.HTMLElement) {
		buyLink := h.ChildAttr("a[class='category-product-item-title-link']", "href")

		c.Visit(buyLink)
	})

	c.OnHTML("div[class='product-page-container']", func(h *colly.HTMLElement) {
		r, _ := regexp.Compile("HKD.*")
		gameName := h.ChildText("h1[class='page-title']")
		buyLink := h.Request.URL.String()
		newPrice := r.FindString(h.ChildText("span[class='special-price']"))
		oldPrice := r.FindString(h.ChildText("span[class='old-price']"))
		times := strings.Split(h.ChildText("span[class='eshop-price-wrapper']"), "\n")
		description := h.ChildText("div[class='value'] > p")
		language := h.ChildText("div[class='product-attribute-val']")

		nintendo := modul.Nintendo{
			GameName:    gameName,
			BuyLink:     buyLink,
			NewPrice:    newPrice,
			OldPrice:    oldPrice,
			Time:        times[0],
			Description: description,
			Language:    language,
		}
		saveDatabase(&nintendo, &i)
	})

	c.Visit("https://store.nintendo.com.hk/games/sale")
	fmt.Print("Price query completed!")
}
