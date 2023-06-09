package main

import (
	"fmt"
	"github.com/Coolenov/Fusion-api-collector/internal"
	"github.com/Coolenov/Fusion-library/database"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	runservice()
}

func runservice() {
	for {
		db := database.DbConnect("root:firstpass@tcp(db:3306)/Fusion_db?utf8mb4&loc=Local")
		links, err := database.GetScrapersUrl(db)
		if err != nil {
			fmt.Println("Cant get scrapers URL!!!\n Trying more...\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, link := range links {
			fmt.Println("Try request ", link)
			internal.GetAndSaveScrapersPosts(link, db)
			err := database.ChangeLastRequestByLink(link, db)
			fmt.Println("Request finished", link)
			if err != nil {
				fmt.Println("Cant change last_request", err)
			}
		}
		err = db.Close()
		if err != nil {
			fmt.Println("Can't close the connection")
		}
		time.Sleep(20 * time.Second)
	}
}
