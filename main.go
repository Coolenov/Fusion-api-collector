package main

import (
	"fmt"
	"github.com/Coolenov/Fusion-api-collector/internal"
	"github.com/Coolenov/Fusion-library/database"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	//dbUrl := os.Getenv("DB_URL")

	runservice()

}

func runservice() {
	for {
		db := database.DbConnect("root:firstpass@tcp(fusiondb:3306)/Fusion_db?utf8mb4&loc=Local")
		links, err := database.GetScrapersUrl(db)
		if err != nil {
			fmt.Println("Cant get scrapers URL!!!\n Trying more...\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, link := range links {
			internal.GetAndSaveScrapersPosts(link, db)
			err := database.ChangeLastRequestByLink(link, db)
			if err != nil {
				fmt.Println("Cant change last_request", err)
			}
		}
		defer db.Close()
		time.Sleep(10 * time.Second)
	}
}
