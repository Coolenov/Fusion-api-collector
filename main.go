package main

import (
	"database/sql"
	"fmt"
	"github.com/Coolenov/Fusion-api-collector/internal"
	"github.com/Coolenov/Fusion-library/database"

	_ "github.com/go-sql-driver/mysql"
	"time"
)

//func init() {
//
//}

func main() {
	dbUrl := "root:root@tcp(db:3306)/Fusion_db?utf8mb4&loc=Local"
	//dbUrl := "root:root@tcp(db:3306)/Fusion_db?utf8mb4"
	//dbUrl := "root:root@tcp(db:3306)/Fusion_db?charset=utf8mb4&collation=utf8mb4_unicode_ci"
	var db *sql.DB
	db = database.DbConnect(dbUrl)
	//database.DBConnect(dbUrl)

	runservice(db)
	defer db.Close()
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

func runservice(db *sql.DB) {
	for {
		links, err := database.GetScrapersUrl(db)
		if err != nil {
			fmt.Println("Cant get scrapers URL!!!\n Trying more...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println(links)
		for _, link := range links {
			internal.GetAndSaveScrapersPosts(link, db)
			err := database.ChangeLastRequestByLink(link, db)
			if err != nil {
				fmt.Println("Cant change last_request", err)
			}
		}
		fmt.Println("for is finished")
		time.Sleep(10 * time.Second)
	}

}
