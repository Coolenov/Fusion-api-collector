package main

import (
	"FusionAPI/initialize"
	"FusionAPI/lib/database"
	"FusionAPI/services/apiDataCollector"
	"fmt"
	"time"
)

func init() {
	initialize.LoadEnv()
	initialize.DBConnect()
}

func main() {

	for {
		links, err := database.GetScrapersUrl()
		if err != nil {
			fmt.Println("Cant get scrapers URL!!!\n Trying more...", err)
			continue
		}
		fmt.Println(links)
		for _, link := range links {
			apiDataCollector.GetAndSaveScrapersPosts(link)
			err := database.ChangeLastRequestByLink(link)
			if err != nil {
				fmt.Println("Cant change last_request", err)
			}
		}
		fmt.Println("for is finished")
		time.Sleep(10 * time.Second)
	}

}
