package main

import (
	"api-collector/internal"
	"fmt"
	"github.com/Coolenov/Fusion-library/database"
	"time"
)

//func init() {
//	initialize.LoadEnv()
//	database.DBConnect()
//}

func main() {
	dbUrl := "root:root@tcp(127.0.0.1:3307)/Fusion_db?charset=utf8mb4&loc=Local"
	database.DBConnect(dbUrl)
	defer database.DB.Close()
	runservice()
	//for {
	//	links, err := database.GetScrapersUrl(database.DB)
	//	if err != nil {
	//		fmt.Println("Cant get scrapers URL!!!\n Trying more...", err)
	//		continue
	//	}
	//	fmt.Println(links)
	//	for _, link := range links {
	//		internal.GetAndSaveScrapersPosts(link, database.DB)
	//		err := database.ChangeLastRequestByLink(link, database.DB)
	//		if err != nil {
	//			fmt.Println("Cant change last_request", err)
	//		}
	//	}
	//	fmt.Println("for is finished")
	//	time.Sleep(10 * time.Second)
	//}
}

func runservice() {
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
