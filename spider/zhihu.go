package spider

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

func Zhihu(wg *sync.WaitGroup) {
	defer wg.Done()
	now := time.Now()
	time := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.UTC).String()
	fileName := strings.Split(time, " ")
	f, CreateErr := os.Create(`./result/zhihu/` + fileName[0] + ".txt")
	if CreateErr != nil {
		fmt.Printf("err: %v\n", CreateErr)
		return
	}
	c := colly.NewCollector()

	c.OnHTML(".HotList-list", func(h *colly.HTMLElement) {
		h.ForEach("section[class='HotItem']", func(i int, h *colly.HTMLElement) {
			title := h.ChildText("h2[class='HotItem-title']")
			content := h.ChildText("p[class='HotItem-excerpt']")

			if title != "" && content != "" {
				_, err1 := f.WriteString("标题: " + title + "\n")
				_, err2 := f.WriteString("内容：" + content + "\n")
				_, err3 := f.WriteString("\n")

				if err1 != nil || err2 != nil || err3 != nil {
					fmt.Printf("写入文件失败!")
				}
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36")
		r.Headers.Set("cookie", `_zap=a4d53e2f-ce2f-4a85-ab1b-1096920c02a2; d_c0="AJDR2dTsvBSPTu0shBvtCigIVQNeHd7NTRI=|1649055934"; _9755xjdesxxd_=32; YD00517437729195%3AWM_TID=%2Fmlyrjim2MFFEFBAEUdqv69Kfxu7TgIb; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1651037657,1651040247,1651040754,1651047937; capsion_ticket=2|1:0|10:1651650435|14:capsion_ticket|44:NzY0ODc2MDBhZDUzNDczYWFiZWY2ZmEzMWI2ODA0YjQ=|244b2066255883d218ac6629c42b7e5cc48111726f5632a541e6eaa15d821ae3; __snaker__id=xnfUpb7vrxoFzYau; YD00517437729195%3AWM_NI=FkcC3CLnlfzXd8h92RA%2Bfynveok8jBjhXpdwveSyqUxJkA2nYG1WBP4or1RXhRe4M%2Bn%2BS5J6AcWNlOccQG5z13pEZYIQQObwr2Slb4Td6ABlMHKopZV0ft3J3%2Fs%2BkREaNHE%3D; YD00517437729195%3AWM_NIKE=9ca17ae2e6ffcda170e2e6ee9bd033a99c8facca65fbac8ea3c14f869b8f87d45e8bba8ad7e45df39d8bd9b62af0fea7c3b92a86b5ada4b24990b2a487f954a8e8a8dab15a83f5aa9aed7ea7b1a5d9e94aaab89891e77db6bb9d9af45fbc96a5a2b1669b99ffb8c15cf8b1bed2d154ac8eabd5d021aef59dd6c93e83bd99afef45a1bd9fa8b872f3edc0a6d9538b8f9fb0e17db2ee9d8de14098bcfaa2c1688f988590f35296f0aa95b242f1f1bdccea5a87bf82a9c837e2a3; q_c1=cac577331f564466bcc8d3ccd97aeb8d|1651833145000|1651833145000; _xsrf=jOB0kdN9FWEEYSGlKcKCQCtzYTgldFaw; NOT_UNREGISTER_WAITING=1; SESSIONID=QWhtRvhlB8GgKLppbYyiJMxm6mdSkBmJ2PzhNSEGQdM; JOID=VFoXBEK0HQzxeH18VLX2WACkkMtE4UphiyUPFD35IUuTPT4IJExccpJzeXlfYpvb71ZUUnd_rJ_BJ_S-LRZ-HfI=; osd=UlodAk6yHQb3dHt8XrP6XgCulsdC4UBnhyMPHjv1J0uZOzIOJEZafpRzc39TZJvR6VpSUn15oJnBLfKyKxZ0G_4=; captcha_session_v2=2|1:0|10:1651852659|18:captcha_session_v2|88:QTRVOFR4Q3NXVkduYlhORTRBbUxDT0VjZlVmbnZrRi9palNyY0NUMWNBeGRBZWdxVklNVUVaZVNtRDVhUVlhRg==|8e69af60d31f413ed3c204424cc215795f87913d2e6cb68c148f7931d25a884c; gdxidpyhxdE=4%2FDr8dOixvzVwK%5CAgb4hxBGdC%5CWkaTpPslpcunVweKrsO7pzQy3eJ%5CSNSowmA%5CSnKtREmad8fXzp8BvOEvmajC7wNVl8cBs83VWWOLZxPTZ90ognUtVwOVBqkuxgkhk17o62l1ouUjDgkgNakEC%5CASSbOene8VuLk0cRIp9eacoOJntq%3A1651853559378; captcha_ticket_v2=2|1:0|10:1651852684|17:captcha_ticket_v2|704:eyJ2YWxpZGF0ZSI6IkNOMzFfQ2lTMVNiaXhWSHlUUWlQUGFSSlBvaS5Jc0xqeGtzaWNvRG90UlRTVW5CWENqc1RucnpJWGMxcG5xRDZxMHZrbjdGUFRIYkJGTVgwRzFJTkY3U2RENzU2T2JrMHlwYzJueFVHYXRfbHJkaUt1S1Z2SzJiSmt0RDV1aUp3bS1aTzEtN0RLMXljbWlNcHl6czJRYURzOUxGOFFTck1NVUZHRGdSQkwwYWsudk9XcUlfQmlTbVQxVUtzTGpldFlPQXZfaWJxNElNXzR3cHJFN3J6YUtRczhVS09EbzQtVG1YTjhXamk1dnVjWnRfeXRMd3N3bUk2dm1MVm5iYWExUXo1enlCUDV6aUxoaEYwekFrREx4ZkR5LVdOMmltMjc1MHdmbmdOZ29VUmpNRm1hZnpqZl9tUXc4aXUxZFhDb0phLmVVSnhMa0laOXJWVXlMV2Q3aUFIRUVkUV9WSHVpZEFtSnJKbS5JSmdNcFdqRU9qZjhWOGM2akh5Nml2VnRTa0ZUUEctSjhDSlBIOTd2SjJKMHlMc0I1UXZCSlpCLmp0dXFZU3dsQ2JjLUJzLV9kZFNKVkpfWkJHZlZBUmNjVGc4OGpzSDRYNGVBY3JtT3ZXeXBQVnN3NDFiQjFjZ2JTbWIwdjVNR2U5X0xCVkdVOXYwNXVkX3JQNEN1bW16MyJ9|002249b6fd19ec1cc625343f7eefe36549917e6d5a0badd9a6c30df72f455e14; z_c0=2|1:0|10:1651852685|4:z_c0|92:Mi4xUy1pN0ZnQUFBQUFBa05IWjFPeThGQ1lBQUFCZ0FsVk5qWk5pWXdBR2xBS3ZnVDd4NDNiVGVXTGpWQzUxTUoxWjNn|1630e6062925656ffa010b62641ac1faa41a59f3ec4697586a6a4c3ee4ca12ac; tst=h; ariaDefaultTheme=undefined; KLBRSID=81978cf28cf03c58e07f705c156aa833|1651854073|1651852536`)
	})

	defer f.Close()
	defer fmt.Println("知乎热搜爬取完成!")

	c.Visit("https://www.zhihu.com/hot")
}
