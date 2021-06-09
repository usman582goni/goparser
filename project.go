package main
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/PuerkitoBio/goquery"
)
func ExampleScrape() {
	for i := 2; i < 3; i++{
		arg := strconv.FormatInt(int64(i), 10)
		res, err := http.Get("https://www.bikester.es/bicicletas.html?page="+arg)

		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".gallery_item").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the title
				href1, _ := s.Find(".js-galleryProductLink").Attr("href")
				//fmt.Printf("Review %d: %s\n", i, href1)

				res1, err1 := http.Get(href1)

				if err1 != nil {
					log.Fatal(err1)
				}
				defer res1.Body.Close()
				if res1.StatusCode != 200 {
					log.Fatalf("status code error: %d %s", res1.StatusCode, res1.Status)
				}
				doc1, err1 := goquery.NewDocumentFromReader(res1.Body)
				if err1 != nil {
					log.Fatal(err1)
				}

				title := doc1.Find("h1.cyc-typo_display-3").Text()
				//tipo := doc1.Find("span.cyc-color-text_secondary").Text()
				tipos := doc1.Find(".pdp_features .js-pdpFeatures .pdp_features--animate").Text()
				fmt.Printf("Title - %s, Tipo - %s", title, tipos)

				//Селект размеров
				doc1.Find(".variation").Each(func(i int, s *goquery.Selection) {
					//attributes := doc1.Find(".variation").Text()
					//fmt.Printf("Attributes %s\n", attributes)
				})

				doc1.Find(".cyc-color-text_secondary").Each(func(i int, s *goquery.Selection) {
					characteristicas := doc1.Find(".cyc-color-text_secondary").Text()
					fmt.Printf("Attributes %s\n", characteristicas)
				})

				fmt.Printf("**************************\n\n")



		})

	}
}
func main() {
	ExampleScrape()
}