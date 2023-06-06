package main

import (
	"api-collector/internal"
	"fmt"
	"time"

	"github.com/Coolenov/Fusion-library/database"
)

//func init() {
//	initialize.LoadEnv()
//	database.DBConnect()
//}

func main() {
	database.DBConnect()
	defer database.DB.Close()
	for {
		links, err := database.GetScrapersUrl(database.DB)
		if err != nil {
			fmt.Println("Cant get scrapers URL!!!\n Trying more...", err)
			continue
		}
		fmt.Println(links)
		for _, link := range links {
			internal.GetAndSaveScrapersPosts(link, database.DB)
			err := database.ChangeLastRequestByLink(link, database.DB)
			if err != nil {
				fmt.Println("Cant change last_request", err)
			}
		}
		fmt.Println("for is finished")
		time.Sleep(10 * time.Second)
	}

}
