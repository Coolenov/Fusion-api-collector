package main

import (
	"database/sql"
	"fmt"
	"github.com/Coolenov/Fusion-api-collector/internal"
	"github.com/Coolenov/Fusion-library/database"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	fmt.Println("FIRST PRINT")
	//dbUrl := os.Getenv("DB_URL")
	db := database.DbConnect("root:firstpass@tcp(database:3306)/Fusion_db?utf8mb4&loc=Local")

	runservice(db)
	defer db.Close()
}

func runservice(db *sql.DB) {
	for {
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
		time.Sleep(10 * time.Second)
	}
}
