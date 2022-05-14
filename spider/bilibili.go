package spider

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

func Bilibili(wg *sync.WaitGroup) {
	defer wg.Done()
	fileName := `./result/bilibili/` + time.Now().Format("20220417") + ".txt"
	hotFan, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer hotFan.Close()

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39")
	})

	c.OnHTML("li[class='rank-item']", func(h *colly.HTMLElement) {
		FanDramaUrl := "https:" + h.ChildAttr("a[class='title']", "href")
		FanDramaName := h.ChildText("a[class='title']")

		_, err2 := hotFan.WriteString("番剧地址: " + FanDramaUrl + "\n")
		_, err3 := hotFan.WriteString("番剧名字: " + FanDramaName + "\n")

		if err2 != nil || err3 != nil {
			fmt.Println("地址或剧名写入失败:")
			fmt.Printf("err2: %v\n", err2)
			fmt.Printf("err3: %v\n", err3)
		}

		c.Visit(FanDramaUrl)
	})

	c.OnHTML("span[class='absolute']", func(h *colly.HTMLElement) {
		detail := h.Text
		_, err2 := hotFan.WriteString("简介: \n")
		_, err3 := hotFan.WriteString(detail + "\n")
		_, err4 := hotFan.WriteString("\n")

		if err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("简介写入失败!")
			fmt.Printf("err2: %v\n", err2)
			fmt.Printf("err3: %v\n", err3)
			fmt.Printf("err4: %v\n", err4)
		}
	})

	defer fmt.Println("番剧爬取完成运行完成!")
	c.Visit("https://www.bilibili.com/v/popular/rank/bangumi")
}
