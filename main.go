package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)

	var items []item

	c.OnHTML("div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		name := h.ChildText("div[itemprop=itemListElement] h2[itemprop=name]")
		price := h.ChildText("div[itemprop=itemListElement] div.sale-price")
		imgUrl := h.ChildAttr("div[itemprop=itemListElement] img", "src")

		if price == "" {
			price = "N/A"
		}
		if imgUrl == "" {
			imgUrl = "N/A"
		}

		if name != "" {
			item := item{Name: name, Price: price, ImgUrl: imgUrl}
			items = append(items, item)
		}
	})

	c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
		next_page := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://j2store.net/demo/index.php/shop")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("products.json", content, 0644)
}
