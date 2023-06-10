package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Coolenov/Fusion-library/database"
	"github.com/Coolenov/Fusion-library/types"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetAndSaveScrapersPosts(link string, db *sql.DB) {
	posts := getScrapedData(link)
	saveScraperPosts(posts, db)
}

func getScrapedData(scraper_link string) []types.Post {
	var posts []types.Post
	client := &http.Client{}

	req, err := http.NewRequest("GET", scraper_link, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return posts
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return posts
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return posts
	}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return posts
	}
	resp.Body.Close()
	return posts
}

func saveScraperPosts(posts []types.Post, db *sql.DB) {

	for _, post := range posts {
		if !database.CheckPostExistByLink(post.Link, db) {
			postId := database.AddPostIntoPostsTable(post, db)
			tags := removeDuplicates(post.Tags)
			for _, tag := range tags {
				var tagId int64
				if !database.CheckTagExist(tag, db) {
					tagId = database.AddTagIntoTagsTable(tag, db)
				} else {
					tagId = database.GetTagIdByTag(tag, db)
				}
				database.AddIntoPostTagsTable(postId, tagId, db)
			}
		}
	}
}

func removeDuplicates(arr []string) []string {
	uniqueMap := make(map[string]bool)
	result := []string{}

	for _, str := range arr {
		str = strings.ToLower(str)
		if !uniqueMap[str] {
			result = append(result, str)
			uniqueMap[str] = true
		}
	}
	return result
}

//func GetAndSavePosts(scraper_link string) {
//	var posts []lib.Post
//	posts = getScrapedData(scraper_link)
//	err := sp(posts)
//	if err != nil {
//		fmt.Println(err)
//	}
//	//savePosts(posts)
//}

//func GetAndSavePosts(scraper_link string) {
//	var posts []lib.Post
//	posts = getScrapedData(scraper_link)
//	for _, post := range posts {
//		err := gormDb.AddPostWithTags(post)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//}

//func savePosts(posts []lib.Post) {
//	//gormDb.AddPostsToDataBase(posts)
//	//gormDb.AddTagsToDataBase(posts)
//	for _, post := range posts {
//		result := initialize.DB.Create(&post)
//		if result.Error != nil {
//			fmt.Println("Ошибка при добавлении поста:", result.Error)
//		}
//	}
//}

//func sp(posts []lib.Post) error {
//	for i := range posts {
//		post := &posts[i]
//		// Получаем или создаем теги и связываем их с постом
//		for j := range post.Tags {
//			tag := &post.Tags[j]
//			if err := initialize.DB.Where("text = ?", tag.Text).FirstOrCreate(tag).Error; err != nil {
//				return err
//			}
//		}
//
//		// Сохраняем пост в базу данных
//		if err := initialize.DB.Create(post).Error; err != nil {
//			return err
//		}
//	}
//	return nil
//}
