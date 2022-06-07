package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Auther struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Auther   Auther    `json:"author"`
	Comments []Comment `json:"comments"`
}

func main() {

	// jsonファイルオープン
	jsonFile, err := os.Open("post.json")
	if err != nil {
		log.Println("Error opening Json file: ", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	for {
		var post Post

		// デコード
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error decoding JSON: ", err)
			return
		}
		fmt.Println(post)
	}

}
